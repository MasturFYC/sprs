package routers

import (
	"fyc.com/sprs/middleware"

	"github.com/gorilla/mux"
)

func LentRouter(router *mux.Router) {

	router.HandleFunc("", middleware.Lent_GetAll).Methods("GET")
	router.HandleFunc("/{id}", middleware.Lent_GetItem).Methods("GET")
	router.HandleFunc("/{id}", middleware.Lent_Delete).Methods("DELETE")
	router.HandleFunc("", middleware.Lent_Create).Methods("POST")
	router.HandleFunc("/{id}", middleware.Lent_Update).Methods("PUT")

}
