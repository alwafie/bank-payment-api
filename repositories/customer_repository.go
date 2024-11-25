package repositories

import (
	"belajar/models"
	"encoding/json"
	"errors"
	"os"
)

const (
	customersFile = "./data/customers.json"
)

func ReadCustomers() ([]models.Customer, error) {
	file, err := os.Open(customersFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var customers []models.Customer
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&customers)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func WriteCustomers(customerId string, newBalance float64) error {
	customers, err := ReadCustomers()
	if err != nil {
		return err
	}

	var customerFound bool
	for i, customer := range customers {
		if customer.ID == customerId {
			customers[i].Balance = newBalance
			customerFound = true
			break
		}
	}

	if !customerFound {
		return errors.New("customer not found")
	}

	file, err := os.Create(customersFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(customers)
	if err != nil {
		return err
	}

	return nil
}
