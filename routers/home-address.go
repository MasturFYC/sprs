package routers

import (
	"fyc.com/sprs/middleware"

	"github.com/gorilla/mux"
)

func HomeAddressRouter(router *mux.Router) {

	router.HandleFunc("/", middleware.CreateHomeAddress).Methods("POST")
	router.HandleFunc("/{id}/", middleware.UpdateHomeAddress).Methods("PUT")
	router.HandleFunc("/{id}/", middleware.DeleteHomeAddress).Methods("DELETE")

}
