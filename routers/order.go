package routers

import (
	"fyc.com/sprs/middleware"

	"github.com/gorilla/mux"
)

func OrderRouter(router *mux.Router) {

	router.HandleFunc("/", middleware.GetOrders).Methods("GET")
	router.HandleFunc("/{id}/", middleware.GetOrder).Methods("GET")
	router.HandleFunc("/{id}/", middleware.DeleteOrder).Methods("DELETE")
	router.HandleFunc("/", middleware.CreateOrder).Methods("POST")
	router.HandleFunc("/{id}/", middleware.UpdateOrder).Methods("PUT")

}
