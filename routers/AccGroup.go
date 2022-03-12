package routers

import (
	"fyc.com/sprs/middleware"

	"github.com/gorilla/mux"
)

func AccGroupRouter(router *mux.Router) {

	router.HandleFunc("/", middleware.GetAccGroups).Methods("GET")
	router.HandleFunc("/all-accounts/", middleware.Group_GetAllAccount).Methods("GET")
	//router.HandleFunc("/props/", middleware.GetTransactionTypeProps).Methods("GET")
	//router.HandleFunc("/{id}/", middleware.GetTransactionType).Methods("GET")
	router.HandleFunc("/{id}/", middleware.DeleteAccGroup).Methods("DELETE")
	router.HandleFunc("/", middleware.CreateAccGroup).Methods("POST")
	router.HandleFunc("/{id}/", middleware.UpdateAccGroup).Methods("PUT")

}
