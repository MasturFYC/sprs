package routers

import (
	"fyc.com/sprs/middleware"

	"github.com/gorilla/mux"
)

func UnitRouter(router *mux.Router) {

	//router.HandleFunc("/", middleware.GetUnit).Methods("GET")
	router.HandleFunc("/{id}/", middleware.GetUnit).Methods("GET")
	router.HandleFunc("/{id}/", middleware.DeleteUnit).Methods("DELETE")
	router.HandleFunc("/", middleware.CreateUnit).Methods("POST")
	router.HandleFunc("/{id}/", middleware.UpdateUnit).Methods("PUT")

}
