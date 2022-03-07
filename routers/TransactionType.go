package routers

import (
	"fyc.com/sprs/middleware"

	"github.com/gorilla/mux"
)

func TransactionTypeRouter(router *mux.Router) {

	router.HandleFunc("/", middleware.GetTransactionTypes).Methods("GET")
	//router.HandleFunc("/props/", middleware.GetTransactionTypeProps).Methods("GET")
	//router.HandleFunc("/{id}/", middleware.GetTransactionType).Methods("GET")
	router.HandleFunc("/{id}/", middleware.DeleteTransactionType).Methods("DELETE")
	router.HandleFunc("/", middleware.CreateTransactionType).Methods("POST")
	router.HandleFunc("/{id}/", middleware.UpdateTransactionType).Methods("PUT")

}
