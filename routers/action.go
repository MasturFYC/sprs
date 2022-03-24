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
	router.HandleFunc("/upload-file/{id}/", middleware.Action_UploadFile).Methods("POST")
	router.HandleFunc("/file/{txt}", middleware.Action_GetFile).Methods("GET")
	router.HandleFunc("/preview/{txt}", middleware.Action_GetPreview).Methods("GET")

}
