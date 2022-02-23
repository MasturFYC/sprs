package routers

import (
	"fyc.com/sprs/middleware"

	"github.com/gorilla/mux"
)

func WheelRouter(router *mux.Router) {

	router.HandleFunc("/", middleware.GetWheels).Methods("GET")
	router.HandleFunc("/{id}/", middleware.GetWheel).Methods("GET")
	router.HandleFunc("/{id}/", middleware.DeleteWheel).Methods("DELETE")
	router.HandleFunc("/", middleware.CreateWheel).Methods("POST")
	router.HandleFunc("/{id}/", middleware.UpdateWheel).Methods("PUT")

}
