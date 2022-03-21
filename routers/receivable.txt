package routers

import (
	"fyc.com/sprs/middleware"

	"github.com/gorilla/mux"
)

func ReceivableRouter(router *mux.Router) {

	//router.HandleFunc("/", middleware.GetReceivable).Methods("GET")
	//router.HandleFunc("/{id}/", middleware.GetReceivable).Methods("GET")
	router.HandleFunc("/{id}/", middleware.DeleteReceivable).Methods("DELETE")
	router.HandleFunc("/", middleware.CreateReceivable).Methods("POST")
	router.HandleFunc("/{id}/", middleware.UpdateReceivable).Methods("PUT")

}
