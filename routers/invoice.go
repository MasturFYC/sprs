package routers

import (
	"fyc.com/sprs/middleware"

	"github.com/gorilla/mux"
)

func InvoiceRouter(router *mux.Router) {

	router.HandleFunc("/", middleware.Invoice_GetAll).Methods("GET")
	router.HandleFunc("/{id}/", middleware.Invoice_GetItem).Methods("GET")
	//	router.HandleFunc("/{id}/", middleware.DeleteFinance).Methods("DELETE")
	//	router.HandleFunc("/", middleware.CreateFinance).Methods("POST")
	//	router.HandleFunc("/{id}/", middleware.UpdateFinance).Methods("PUT")

}
