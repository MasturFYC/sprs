package routers

import (
	"fyc.com/sprs/middleware"

	"github.com/gorilla/mux"
)

func OrderRouter(router *mux.Router) {

	router.HandleFunc("/", middleware.GetOrders).Methods("GET")
	router.HandleFunc("/search/{txt}/", middleware.SearchOrders).Methods("GET")
	router.HandleFunc("/finance/{id}/", middleware.GetOrdersByFinance).Methods("GET")
	router.HandleFunc("/branch/{id}/", middleware.GetOrdersByBranch).Methods("GET")
	router.HandleFunc("/month/{id}/", middleware.GetOrdersByMonth).Methods("GET")
	//router.HandleFunc("/orders/search/", middleware.IsAuthorized(middleware.GetOrders)).Methods("GET")
	router.HandleFunc("/{id}/", middleware.GetOrder).Methods("GET")
	router.HandleFunc("/{id}/", middleware.DeleteOrder).Methods("DELETE")
	//router.HandleFunc("/{id}/", middleware.IsAuthorized(middleware.DeleteOrder)).Methods("DELETE")
	router.HandleFunc("/", middleware.CreateOrder).Methods("POST")
	router.HandleFunc("/{id}/", middleware.UpdateOrder).Methods("PUT")

}
