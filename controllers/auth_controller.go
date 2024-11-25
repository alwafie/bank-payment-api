package controllers

import (
	"belajar/middleware"
	"belajar/services"
	"belajar/utils"
	"encoding/json"
	"net/http"
)

type CustomerController struct {
	service *services.CustomerService
}

func NewCustomerHandler(service *services.CustomerService) *CustomerController {
	return &CustomerController{service: service}
}

func (h *CustomerController) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	customerID, token, err := h.service.Login(req.Username, req.Password)
	if err != nil {
		if err == services.ErrCustomerNotFound {
			utils.RespondWithError(w, http.StatusUnauthorized, "Customer not found")
		} else {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"customer_id": customerID,
		"token":       token,
	})
}

func (h *CustomerController) Logout(w http.ResponseWriter, r *http.Request) {
	token, ok := r.Context().Value(middleware.TokenContextKey).(string)
	if !ok || token == "" {
		utils.RespondWithError(w, http.StatusUnauthorized, "Missing or invalid authorization token")
		return
	}

	err := h.service.Logout(token)
	if err != nil {
		if err == services.ErrNotLoggedIn {
			utils.RespondWithError(w, http.StatusUnauthorized, "Customer is not logged in")
		} else {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Logout successful",
	})
}
