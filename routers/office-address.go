package routers

import (
	"fyc.com/sprs/middleware"

	"github.com/gorilla/mux"
)

func OfficeAddressRouter(router *mux.Router) {

	router.HandleFunc("/", middleware.CreateOfficeAddress).Methods("POST")
	router.HandleFunc("/{id}/", middleware.UpdateOfficeAddress).Methods("PUT")
	router.HandleFunc("/{id}/", middleware.DeleteOfficeAddress).Methods("DELETE")

}
