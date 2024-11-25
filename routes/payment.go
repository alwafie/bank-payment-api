package routes

import (
	"belajar/controllers"
	"belajar/middleware"
	"belajar/services"

	"github.com/gorilla/mux"
)

func PaymentRoutes(r *mux.Router) {
	router := r.PathPrefix("/payment").Subrouter()

	paymentService := services.NewPaymentService()
	paymentController := controllers.NewPaymentController(paymentService)

	router.Use(middleware.TokenAuthMiddleware)

	router.HandleFunc("/make", paymentController.MakePayment).Methods("POST")
}
