package routers

import (
	"fyc.com/sprs/middleware"

	"github.com/gorilla/mux"
)

func PostAddressRouter(router *mux.Router) {

	router.HandleFunc("/", middleware.CreatePostAddress).Methods("POST")
	router.HandleFunc("/{id}/", middleware.UpdatePostAddress).Methods("PUT")
	router.HandleFunc("/{id}/", middleware.DeletePostAddress).Methods("DELETE")

}
