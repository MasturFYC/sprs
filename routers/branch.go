package routers

import (
	"fyc.com/sprs/middleware"

	"github.com/gorilla/mux"
)

func BranchRouter(router *mux.Router) {

	router.HandleFunc("", middleware.GetBranchs).Methods("GET")
	router.HandleFunc("/{id}", middleware.GetBranch).Methods("GET")
	router.HandleFunc("/{id}", middleware.DeleteBranch).Methods("DELETE")
	router.HandleFunc("", middleware.CreateBranch).Methods("POST")
	router.HandleFunc("/{id}", middleware.UpdateBranch).Methods("PUT")

}
