package routers

import (
	"fyc.com/sprs/middleware"

	"github.com/gorilla/mux"
)

func TypeRouter(router *mux.Router) {

	router.HandleFunc("/", middleware.GetTypes).Methods("GET")
	router.HandleFunc("/{id}/", middleware.GetType).Methods("GET")
	router.HandleFunc("/{id}/", middleware.DeleteType).Methods("DELETE")
	router.HandleFunc("/", middleware.CreateType).Methods("POST")
	router.HandleFunc("/{id}/", middleware.UpdateType).Methods("PUT")

}
