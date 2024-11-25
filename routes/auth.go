package routes

import (
	"belajar/controllers"
	"belajar/middleware"
	"belajar/services"
	"net/http"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	router := r.PathPrefix("/auth").Subrouter()

	customerService := services.NewCustomerService()

	customerController := controllers.NewCustomerHandler(customerService)

	router.HandleFunc("/login", customerController.Login).Methods("POST")

	router.HandleFunc("/logout", customerController.Logout).Methods("POST").
		Handler(middleware.TokenAuthMiddleware(http.HandlerFunc(customerController.Logout)))
}
