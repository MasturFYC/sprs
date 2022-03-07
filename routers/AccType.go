package routers

import (
	"fyc.com/sprs/middleware"

	"github.com/gorilla/mux"
)

func AccountTypeRouter(router *mux.Router) {

	router.HandleFunc("/", middleware.GetAccountTypes).Methods("GET")
	//router.HandleFunc("/{id}/", middleware.GetAccountType).Methods("GET")
	router.HandleFunc("/{id}/", middleware.DeleteAccountType).Methods("DELETE")
	router.HandleFunc("/", middleware.CreateAccountType).Methods("POST")
	router.HandleFunc("/{id}/", middleware.UpdateAccountType).Methods("PUT")

}
