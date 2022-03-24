package routers

import (
	"fyc.com/sprs/middleware"

	"github.com/gorilla/mux"
)

func OrderRouter(router *mux.Router) {

	router.HandleFunc("", middleware.GetOrders).Methods("GET")
	router.HandleFunc("/search", middleware.SearchOrders).Methods("POST")
	router.HandleFunc("/finance/{id}", middleware.GetOrdersByFinance).Methods("GET")
	router.HandleFunc("/branch/{id}", middleware.GetOrdersByBranch).Methods("GET")
	router.HandleFunc("/month/{id}", middleware.GetOrdersByMonth).Methods("GET")
	//router.HandleFunc("/orders/search/", middleware.IsAuthorized(middleware.GetOrders)).Methods("GET")
	router.HandleFunc("/{id}", middleware.GetOrder).Methods("GET")
	router.HandleFunc("/{id}", middleware.DeleteOrder).Methods("DELETE")
	//router.HandleFunc("/{id}/", middleware.IsAuthorized(middleware.DeleteOrder)).Methods("DELETE")
	router.HandleFunc("", middleware.CreateOrder).Methods("POST")
	router.HandleFunc("/{id}", middleware.UpdateOrder).Methods("PUT")
	router.HandleFunc("/name-seq", middleware.Order_GetNameSeq).Methods("GET")
	router.HandleFunc("/invoiced/{month}/{year}/{financeId}", middleware.Order_GetInvoiced).Methods("GET")

}
