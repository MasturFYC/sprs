package routers

import (
	"fyc.com/sprs/middleware"

	"github.com/gorilla/mux"
)

func LoanRouter(router *mux.Router) {

	router.HandleFunc("", middleware.Loan_GetAll).Methods("GET")
	router.HandleFunc("/{id}", middleware.Loan_GetItem).Methods("GET")
	router.HandleFunc("/{id}", middleware.Loan_Delete).Methods("DELETE")
	router.HandleFunc("", middleware.Loan_Create).Methods("POST")
	router.HandleFunc("/payment/{id}", middleware.Loan_Payment).Methods("POST")
	router.HandleFunc("/{id}", middleware.Loan_Update).Methods("PUT")

}
