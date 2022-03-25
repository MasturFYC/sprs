package routers

import (
	"fyc.com/sprs/middleware"

	"github.com/gorilla/mux"
)

func CustomerRouter(router *mux.Router) {
	// router.HandleFunc("/", middleware.GetCustomers).Methods("GET")

	// id = order id
	router.HandleFunc("", middleware.CreateCustomer).Methods("POST")
	router.HandleFunc("/{id}", middleware.GetCustomer).Methods("GET")
	router.HandleFunc("/{id}", middleware.DeleteCustomer).Methods("DELETE")
	router.HandleFunc("/{id}", middleware.UpdateCustomer).Methods("PUT")

}
