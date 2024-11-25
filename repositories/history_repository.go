package repositories

import (
	"belajar/models"
	"encoding/json"
	"os"
	"time"
)

const (
	historyFile = "./data/history.json"
)

func WriteHistoryAuth(action string, customer models.Customer) error {
	history := models.HistoryAuth{
		Action:       action,
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		Timestamp:    time.Now(),
	}

	file, err := os.OpenFile(historyFile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(history)
}

func WriteHistoryPayment(status string, customer *models.Customer, merchant *models.Merchant) error {
	history := models.HistoryPayment{
		Action:                "payment",
		CustomerID:            customer.ID,
		CustomerName:          customer.Name,
		CustomerAccountNumber: customer.AccountNumber,
		Merchant:              merchant.Name,
		MerchantAccountNumber: merchant.AccountNumber,
		Timestamp:             time.Now(),
		Status:                status,
	}

	file, err := os.OpenFile(historyFile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(history)
}
