package services

import (
	"belajar/models"
	"belajar/repositories"
	"belajar/utils"
	"errors"
)

var ErrCustomerNotLoggedIn = errors.New("customer not logged in")
var ErrSenderNotFound = errors.New("sender customer not found")
var ErrRecipientNotFound = errors.New("recipient merchant not found")
var ErrInsufficientBalance = errors.New("insufficient balance")

type PaymentService struct{}

func NewPaymentService() *PaymentService {
	return &PaymentService{}
}

func (s *PaymentService) MakePayment(tokenString, merchantNumberAccount string, amount float64) (models.Customer, models.Merchant, string, error) {
	status := "success"

	customerClaims, err := utils.ValidateToken(tokenString)
	if err != nil {
		status = "failed"
		_ = repositories.WriteHistoryPayment(status, nil, nil)
		return models.Customer{}, models.Merchant{}, status, ErrCustomerNotLoggedIn
	}

	customerID := customerClaims.(*utils.CustomClaims).ID

	customers, err := repositories.ReadCustomers()
	if err != nil {
		status = "failed"
		_ = repositories.WriteHistoryPayment(status, nil, nil)
		return models.Customer{}, models.Merchant{}, status, err
	}

	var customer *models.Customer
	for _, c := range customers {
		if c.ID == customerID {
			customer = &c
			break
		}
	}

	if customer == nil {
		status = "failed"
		_ = repositories.WriteHistoryPayment(status, nil, nil)
		return models.Customer{}, models.Merchant{}, status, ErrSenderNotFound
	}

	merchants, err := repositories.ReadMerchants()
	if err != nil {
		status = "failed"
		_ = repositories.WriteHistoryPayment(status, customer, nil)
		return models.Customer{}, models.Merchant{}, status, err
	}

	var merchant *models.Merchant
	for _, m := range merchants {
		if m.AccountNumber == merchantNumberAccount {
			merchant = &m
			break
		}
	}

	if merchant == nil {
		status = "failed"
		_ = repositories.WriteHistoryPayment(status, customer, nil)
		return models.Customer{}, models.Merchant{}, status, ErrRecipientNotFound
	}

	if customer.Balance < amount {
		status = "failed"
		_ = repositories.WriteHistoryPayment(status, customer, merchant)
		return models.Customer{}, models.Merchant{}, status, ErrInsufficientBalance
	}

	customer.Balance -= amount
	merchant.Balance += amount

	err = repositories.WriteCustomers(customer.ID, customer.Balance)
	if err != nil {
		status = "failed"
		_ = repositories.WriteHistoryPayment(status, customer, merchant)
		return models.Customer{}, models.Merchant{}, status, err
	}

	err = repositories.WriteMerchants(merchantNumberAccount, merchant.Balance)
	if err != nil {
		status = "failed"
		_ = repositories.WriteHistoryPayment(status, customer, merchant)
		return models.Customer{}, models.Merchant{}, status, err
	}

	err = repositories.WriteHistoryPayment("payment", customer, merchant)
	if err != nil {
		status = "failed"
		_ = repositories.WriteHistoryPayment(status, customer, merchant)
		return models.Customer{}, models.Merchant{}, status, err
	}

	return *customer, *merchant, status, nil
}
