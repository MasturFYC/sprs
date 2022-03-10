package routers

import (
	"fyc.com/sprs/middleware"

	"github.com/gorilla/mux"
)

func ReportRouter(router *mux.Router) {

	router.HandleFunc("/trx/month/{month}/{year}/", middleware.GetRepotTrxByMonth).Methods("GET")
	router.HandleFunc("/trx/month/acc/{acc}/{month}/{year}/", middleware.GetRepotTrxByAccMonth).Methods("GET")

}
