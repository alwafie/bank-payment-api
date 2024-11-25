package controllers

import (
	"belajar/middleware"
	"belajar/services"
	"belajar/utils"
	"encoding/json"
	"net/http"
)

type PaymentController struct {
	service *services.PaymentService
}

func NewPaymentController(service *services.PaymentService) *PaymentController {
	return &PaymentController{service: service}
}

func (h *PaymentController) MakePayment(w http.ResponseWriter, r *http.Request) {
	var req struct {
		MerchantAccount string  `json:"merchantAccountNumber"`
		Amount          float64 `json:"amount"`
	}

	token, ok := r.Context().Value(middleware.TokenContextKey).(string)
	if !ok || token == "" {
		utils.RespondWithError(w, http.StatusUnauthorized, "Missing or invalid authorization token")
		return
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	customer, merchant, status, err := h.service.MakePayment(token, req.MerchantAccount, req.Amount)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"customerId":            customer.ID,
		"customerName":          customer.Name,
		"merchantName":          merchant.Name,
		"merchantAccountNumber": merchant.AccountNumber,
		"amount":                req.Amount,
		"status":                status,
	})
}
