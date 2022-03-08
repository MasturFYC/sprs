package routers

import (
	"fyc.com/sprs/middleware"

	"github.com/gorilla/mux"
)

func SaldoRouter(router *mux.Router) {

	router.HandleFunc("/", middleware.GetRemainSaldo).Methods("GET")
}
