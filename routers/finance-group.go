package routers

import (
	"fyc.com/sprs/middleware"

	"github.com/gorilla/mux"
)

func FinanceGroupRouter(router *mux.Router) {

	router.HandleFunc("", middleware.FinanceGroup_GetAll).Methods("GET")
	router.HandleFunc("/finances", middleware.FinanceGroup_GetFinances).Methods("GET")
	router.HandleFunc("/{id}", middleware.FinanceGroup_GetItem).Methods("GET")
	router.HandleFunc("/{id}", middleware.FinanceGroup_Delete).Methods("DELETE")
	router.HandleFunc("", middleware.FinanceGroup_Create).Methods("POST")
	router.HandleFunc("/{id}", middleware.FinanceGroup_Update).Methods("PUT")

}
