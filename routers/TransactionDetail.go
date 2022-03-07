package routers

import (
	"fyc.com/sprs/middleware"

	"github.com/gorilla/mux"
)

func TransactionDetailRouter(router *mux.Router) {

	router.HandleFunc("/{id}/", middleware.GetTransactionDetails).Methods("GET")
	//router.HandleFunc("/props/", middleware.GetTransactionDetailProps).Methods("GET")
	//router.HandleFunc("/{id}/", middleware.GetTransactionDetail).Methods("GET")
	router.HandleFunc("/{id}/", middleware.DeleteTransactionDetail).Methods("DELETE")
	router.HandleFunc("/", middleware.CreateTransactionDetail).Methods("POST")
	router.HandleFunc("/{id}/", middleware.UpdateTransactionDetail).Methods("PUT")

}
