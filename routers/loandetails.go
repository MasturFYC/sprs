package routers

import (
	"fyc.com/sprs/middleware"

	"github.com/gorilla/mux"
)

func LoanDetailRouter(router *mux.Router) {

	// router.HandleFunc("", middleware.Lent_GetAll).Methods("GET")
	// router.HandleFunc("/{id}", middleware.Lent_GetItem).Methods("GET")
	router.HandleFunc("/{id}", middleware.LoanDetail_Delete).Methods("DELETE")
	router.HandleFunc("", middleware.LoanDetail_Create).Methods("POST")
	router.HandleFunc("/{id}", middleware.LoanDetail_Update).Methods("PUT")

}
