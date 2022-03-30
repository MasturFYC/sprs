package routers

import (
	"fyc.com/sprs/middleware"

	"github.com/gorilla/mux"
)

func LentDetailRouter(router *mux.Router) {

	// router.HandleFunc("", middleware.Lent_GetAll).Methods("GET")
	// router.HandleFunc("/{id}", middleware.Lent_GetItem).Methods("GET")
	router.HandleFunc("/{id}", middleware.LentDetail_Delete).Methods("DELETE")
	router.HandleFunc("", middleware.LentDetail_Create).Methods("POST")
	router.HandleFunc("/{id}", middleware.LentDetail_Update).Methods("PUT")

}
