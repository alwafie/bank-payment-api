package models

import "time"

type HistoryAuth struct {
	Action       string    `json:"action"`
	CustomerID   string    `json:"customerId"`
	CustomerName string    `json:"customerName"`
	Timestamp    time.Time `json:"timestamp"`
}
