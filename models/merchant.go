package models

type Merchant struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	AccountNumber string  `json:"accountNumber"`
	Balance       float64 `json:"balance"`
}
