package routers

import (
	"fyc.com/sprs/middleware"

	"github.com/gorilla/mux"
)

func FinanceRouter(router *mux.Router) {

	router.HandleFunc("/", middleware.GetFinances).Methods("GET")
	router.HandleFunc("/{id}/", middleware.GetFinance).Methods("GET")
	router.HandleFunc("/{id}/", middleware.DeleteFinance).Methods("DELETE")
	router.HandleFunc("/", middleware.CreateFinance).Methods("POST")
	router.HandleFunc("/{id}/", middleware.UpdateFinance).Methods("PUT")

}
