package services

import (
	"belajar/models"
	"belajar/repositories"
	"belajar/utils"
	"errors"
)

var (
	ErrCustomerNotFound = errors.New("customer not found")
	ErrNotLoggedIn      = errors.New("customer is not logged in")
)

type CustomerService struct {
	loggedInCustomers map[string]bool
}

func NewCustomerService() *CustomerService {
	return &CustomerService{loggedInCustomers: make(map[string]bool)}
}

func (cs *CustomerService) Login(username, password string) (string, string, error) {
	customers, err := repositories.ReadCustomers()
	if err != nil {
		return "", "", err
	}

	for _, customer := range customers {
		if customer.Username == username {
			if err := utils.VerifyPassword(customer.Password, password); err != nil {
				return "", "", err
			}

			cs.loggedInCustomers[customer.ID] = true

			if err := repositories.WriteHistoryAuth("login", customer); err != nil {
				return "", "", err
			}

			token, err := utils.CreateToken(&customer)
			if err != nil {
				return "", "", err
			}

			return customer.ID, token, nil
		}
	}

	return "", "", ErrCustomerNotFound
}

func (cs *CustomerService) Logout(token string) error {
	customerClaims, err := utils.ValidateToken(token)
	if err != nil {
		return ErrCustomerNotLoggedIn
	}

	customerID := customerClaims.(*utils.CustomClaims).ID
	if _, ok := cs.loggedInCustomers[customerID]; !ok {
		return ErrNotLoggedIn
	}

	delete(cs.loggedInCustomers, customerID)

	if err := utils.AddToBlacklist(token); err != nil {
		return err
	}

	customers, err := repositories.ReadCustomers()
	if err != nil {
		return err
	}

	var customer *models.Customer
	for _, c := range customers {
		if c.ID == customerID {
			customer = &c
			break
		}
	}

	if customer != nil {
		if err := repositories.WriteHistoryAuth("logout", *customer); err != nil {
			return err
		}
	}

	return nil
}
