package routers

import (
	"fyc.com/sprs/middleware"

	"github.com/gorilla/mux"
)

func TaskRouter(router *mux.Router) {

	router.HandleFunc("", middleware.CreateTask).Methods("POST")
	router.HandleFunc("/{id}", middleware.GetTask).Methods("GET")
	router.HandleFunc("/{id}", middleware.DeleteTask).Methods("DELETE")
	router.HandleFunc("/{id}", middleware.UpdateTask).Methods("PUT")

}
