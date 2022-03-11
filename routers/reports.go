package routers

import (
	"fyc.com/sprs/middleware"

	"github.com/gorilla/mux"
)

func ReportRouter(router *mux.Router) {

	router.HandleFunc("/trx/month/{month}/{year}/", middleware.GetRepotTrxByMonth).Methods("GET")
	//router.HandleFunc("/trx/month/type/{type}/{month}/{year}/", middleware.GetRepotTrxByTypeMonth).Methods("GET")
	//router.HandleFunc("/trx/month/acc/{acc}/{type}/{month}/{year}/", middleware.GetRepotTrxByAccountMonth).Methods("GET")

}
