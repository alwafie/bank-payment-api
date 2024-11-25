package models

type Customer struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	AccountNumber string  `json:"accountNumber"`
	Balance       float64 `json:"balance"`
	Username      string  `json:"username"`
	Password      string  `json:"password"`
}
