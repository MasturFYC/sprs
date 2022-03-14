package routers

import (
	"fyc.com/sprs/middleware"

	"github.com/gorilla/mux"
)

func AccountCodeRouter(router *mux.Router) {

	router.HandleFunc("/", middleware.GetAccountCodes).Methods("GET")
	router.HandleFunc("/search-name/{txt}/", middleware.SearchAccountCodeByName).Methods("GET")
	//router.HandleFunc("/search-name/{txt}/", middleware.SearchAccountCodeByName).Methods("GET")
	router.HandleFunc("/group-type/{id}/", middleware.GetAccountCodeByType).Methods("GET")
	//router.HandleFunc("/props/", middleware.GetAccountCodeProps).Methods("GET")
	router.HandleFunc("/props/", middleware.GetAccountCodeProps).Methods("GET")
	router.HandleFunc("/{id}/", middleware.GetAccountCode).Methods("GET")
	router.HandleFunc("/spec/{id}/", middleware.Account_GetSpec).Methods("GET")
	router.HandleFunc("/{id}/", middleware.DeleteAccountCode).Methods("DELETE")
	router.HandleFunc("/", middleware.CreateAccountCode).Methods("POST")
	router.HandleFunc("/{id}/", middleware.UpdateAccountCode).Methods("PUT")

}
