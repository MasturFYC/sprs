package routers

import (
	"fyc.com/sprs/middleware"

	"github.com/gorilla/mux"
)

func MerkRouter(router *mux.Router) {

	router.HandleFunc("/", middleware.GetMerks).Methods("GET")
	router.HandleFunc("/{id}/", middleware.GetMerk).Methods("GET")
	router.HandleFunc("/{id}/", middleware.DeleteMerk).Methods("DELETE")
	router.HandleFunc("/", middleware.CreateMerk).Methods("POST")
	router.HandleFunc("/{id}/", middleware.UpdateMerk).Methods("PUT")

}
