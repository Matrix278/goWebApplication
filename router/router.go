package router

import (
	"goWebApplication/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	// router.HandleFunc("/customer/{id}", middleware.GetCustomer).Methods(http.MethodGet, "OPTIONS")
	router.HandleFunc("/customers", middleware.GetAllCustomers).Methods(http.MethodGet, "OPTIONS")
	// router.HandleFunc("/customer/create", middleware.CreateCustomer).Methods(http.MethodPost, "OPTIONS")
	// router.HandleFunc("/customer/{id}/update", middleware.UpdateCustomer).Methods(http.MethodPut, "OPTIONS")
	// router.HandleFunc("/customer/{id}/delete", middleware.DeleteCustomer).Methods(http.MethodDelete, "OPTIONS")

	return router
}
