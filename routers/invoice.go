package routers

import (
	"fyc.com/sprs/middleware"

	"github.com/gorilla/mux"
)

func InvoiceRouter(router *mux.Router) {

	router.HandleFunc("/", middleware.Invoice_GetAll).Methods("GET")
	router.HandleFunc("/{invoceId}/", middleware.Invoice_GetItem).Methods("GET")
	router.HandleFunc("/{financeId}/{invoiceId}/", middleware.Invoice_GetOrders).Methods("GET")
	router.HandleFunc("/", middleware.Invoice_Create).Methods("POST")
	router.HandleFunc("/{id}/", middleware.Invoice_Update).Methods("PUT")

}
