package routers

import (
	"fyc.com/sprs/middleware"

	"github.com/gorilla/mux"
)

func TransactionRouter(router *mux.Router) {

	router.HandleFunc("/", middleware.GetTransactions).Methods("GET")
	router.HandleFunc("/search/", middleware.SearchTransactions).Methods("POST")
	router.HandleFunc("/group-type/{id}/", middleware.GetTransactionsByType).Methods("GET")
	router.HandleFunc("/month/{id}/", middleware.GetTransactionsByMonth).Methods("GET")
	//router.HandleFunc("/props/", middleware.GetTransactionProps).Methods("GET")
	//router.HandleFunc("/{id}/", middleware.GetTransaction).Methods("GET")
	router.HandleFunc("/{id}/", middleware.DeleteTransaction).Methods("DELETE")
	router.HandleFunc("/", middleware.CreateTransaction).Methods("POST")
	router.HandleFunc("/{id}/", middleware.UpdateTransaction).Methods("PUT")

}
