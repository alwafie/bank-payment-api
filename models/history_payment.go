package models

import "time"

type HistoryPayment struct {
	Action                string    `json:"action"`
	CustomerID            string    `json:"customerId"`
	CustomerName          string    `json:"customerName"`
	CustomerAccountNumber string    `json:"customerAccountNumber"`
	Merchant              string    `json:"merchantName"`
	MerchantAccountNumber string    `json:"merchantNumberAccount"`
	Timestamp             time.Time `json:"timestamp"`
	Status                string    `json:"status"`
}
