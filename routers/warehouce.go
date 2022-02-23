package routers

import (
	"fyc.com/sprs/middleware"

	"github.com/gorilla/mux"
)

func WarehouseRouter(router *mux.Router) {

	router.HandleFunc("/", middleware.GetWarehouses).Methods("GET")
	router.HandleFunc("/", middleware.CreateWarehouse).Methods("POST")
	router.HandleFunc("/{id}/", middleware.GetWarehouse).Methods("GET")
	router.HandleFunc("/{id}/", middleware.DeleteWarehouse).Methods("DELETE")
	router.HandleFunc("/{id}/", middleware.UpdateWarehouse).Methods("PUT")

}
