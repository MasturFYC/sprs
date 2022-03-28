package routers

import (
	"fyc.com/sprs/middleware"
	"fyc.com/sprs/middleware/reports"

	"github.com/gorilla/mux"
)

func ReportRouter(router *mux.Router) {

	router.HandleFunc("/trx/month/{month}/{year}/", middleware.GetRepotTrxByMonth).Methods("GET")
	//router.HandleFunc("/orders", reports.ReportOrder).Methods("GET")

	// financeId, branchId, typeId, dateFrom, dateTo
	router.HandleFunc("/order-status/{finance}/{branch}/{type}/{month}/{year}/{from}/{to}", reports.ReportOrder).Methods("GET")
	router.HandleFunc("/order-all-waiting/{finance}/{branch}/{type}", reports.ReportOrderAllWaiting).Methods("GET")

	//router.HandleFunc("/trx/month/type/{type}/{month}/{year}/", middleware.GetRepotTrxByTypeMonth).Methods("GET")
	//router.HandleFunc("/trx/month/acc/{acc}/{type}/{month}/{year}/", middleware.GetRepotTrxByAccountMonth).Methods("GET")

}
