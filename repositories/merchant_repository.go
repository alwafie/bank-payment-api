package repositories

import (
	"belajar/models"
	"encoding/json"
	"errors"
	"os"
)

const (
	merchantsFile = "./data/merchants.json"
)

func ReadMerchants() ([]models.Merchant, error) {
	file, err := os.Open(merchantsFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var merchants []models.Merchant
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&merchants)
	if err != nil {
		return nil, err
	}
	return merchants, nil
}

func WriteMerchants(accountNumber string, newBalance float64) error {
	merchants, err := ReadMerchants()
	if err != nil {
		return err
	}

	var merchantFound bool
	for i, merchant := range merchants {
		if merchant.AccountNumber == accountNumber {
			merchants[i].Balance = newBalance
			merchantFound = true
			break
		}
	}

	if !merchantFound {
		return errors.New("merchant not found")
	}

	file, err := os.Create(merchantsFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(merchants)
	if err != nil {
		return err
	}

	return nil
}
