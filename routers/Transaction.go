package routers

import (
	"fyc.com/sprs/middleware"

	"github.com/gorilla/mux"
)

func TransactionRouter(router *mux.Router) {

	router.HandleFunc("/", middleware.GetTransactions).Methods("GET")
	//router.HandleFunc("/props/", middleware.GetTransactionProps).Methods("GET")
	//router.HandleFunc("/{id}/", middleware.GetTransaction).Methods("GET")
	router.HandleFunc("/{id}/", middleware.DeleteTransaction).Methods("DELETE")
	router.HandleFunc("/", middleware.CreateTransaction).Methods("POST")
	router.HandleFunc("/{id}/", middleware.UpdateTransaction).Methods("PUT")

}
