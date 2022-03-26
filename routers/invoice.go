package routers

import (
	"fyc.com/sprs/middleware"

	"github.com/gorilla/mux"
)

func InvoiceRouter(router *mux.Router) {

	router.HandleFunc("/", middleware.Invoice_GetAll).Methods("GET")
	router.HandleFunc("/search/", middleware.Invoice_GetSearch).Methods("POST")
	router.HandleFunc("/month-year/{month}/{year}/", middleware.Invoice_GetByMonth).Methods("GET")
	router.HandleFunc("/finance/{id}/", middleware.Invoice_GetByFinance).Methods("GET")
	router.HandleFunc("/{id}/", middleware.Invoice_GetItem).Methods("GET")
	router.HandleFunc("/{financeId}/{id}/", middleware.Invoice_GetOrders).Methods("GET")
	router.HandleFunc("/", middleware.Invoice_Create).Methods("POST")
	router.HandleFunc("/{id}/", middleware.Invoice_Update).Methods("PUT")
	router.HandleFunc("/{id}/", middleware.Invoice_Delete).Methods("DELETE")
	router.HandleFunc("/download/{id}", middleware.Pdf_GetInvoice).Methods("GET")

	//CLIPAN  : 2
	router.HandleFunc("/download/2/{id}", middleware.Clipan_GetInvoice).Methods("GET")
	// MTF  : 3
	router.HandleFunc("/download/3/{id}", middleware.Mtf_GetInvoice).Methods("GET")
}
