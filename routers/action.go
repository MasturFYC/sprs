package routers

import (
	"fyc.com/sprs/middleware"

	"github.com/gorilla/mux"
)

func ActionRouter(router *mux.Router) {

	router.HandleFunc("/{id}/", middleware.GetActions).Methods("GET")
	//router.HandleFunc("/{id}/", middleware.GetAction).Methods("GET")
	router.HandleFunc("/{id}/", middleware.DeleteAction).Methods("DELETE")
	router.HandleFunc("/", middleware.CreateAction).Methods("POST")
	router.HandleFunc("/{id}/", middleware.UpdateAction).Methods("PUT")

}
