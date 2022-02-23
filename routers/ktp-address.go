package routers

import (
	"fyc.com/sprs/middleware"

	"github.com/gorilla/mux"
)

func KtpAddressRouter(router *mux.Router) {

	router.HandleFunc("/", middleware.CreateKTPAddress).Methods("POST")
	router.HandleFunc("/{id}/", middleware.UpdateKTPAddress).Methods("PUT")
	router.HandleFunc("/{id}/", middleware.DeleteKTPAddress).Methods("DELETE")

}
