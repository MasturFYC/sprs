package main

import (
	"fmt"

	"log"

	"net/http"

	"fyc.com/sprs/routers"

	"github.com/gorilla/mux"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

var (
	mainRouter *mux.Router
)

func createRouter() {
	mainRouter = mux.NewRouter()
}

func loadEnvirontment() {
	err := godotenv.Load("/home/mastur/.env")
	//err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func loadRouter() {

	routers.InitializeRoute(mainRouter)

	routers.AccGroupRouter(mainRouter.PathPrefix("/api/acc-group/").Subrouter())
	routers.AccountTypeRouter(mainRouter.PathPrefix("/api/acc-type/").Subrouter())
	routers.AccountCodeRouter(mainRouter.PathPrefix("/api/acc-code/").Subrouter())
	routers.ActionRouter(mainRouter.PathPrefix("/api/actions/").Subrouter())
	routers.BranchRouter(mainRouter.PathPrefix("/api/branchs/").Subrouter())
	routers.CustomerRouter(mainRouter.PathPrefix("/api/customers/").Subrouter())
	routers.FinanceGroupRouter(mainRouter.PathPrefix("/api/finance-group").Subrouter())
	routers.FinanceRouter(mainRouter.PathPrefix("/api/finances/").Subrouter())
	routers.HomeAddressRouter(mainRouter.PathPrefix("/api/home-address/").Subrouter())
	routers.KtpAddressRouter(mainRouter.PathPrefix("/api/ktp-address/").Subrouter())
	routers.MerkRouter(mainRouter.PathPrefix("/api/merks/").Subrouter())
	routers.OfficeAddressRouter(mainRouter.PathPrefix("/api/office-address/").Subrouter())
	routers.OrderRouter(mainRouter.PathPrefix("/api/orders/").Subrouter())
	routers.PostAddressRouter(mainRouter.PathPrefix("/api/post-address/").Subrouter())
	// routers.ReceivableRouter(mainRouter.PathPrefix("/api/receivables/").Subrouter())
	routers.TaskRouter(mainRouter.PathPrefix("/api/tasks/").Subrouter())
	routers.TypeRouter(mainRouter.PathPrefix("/api/types/").Subrouter())
	routers.UnitRouter(mainRouter.PathPrefix("/api/units/").Subrouter())
	routers.WarehouseRouter(mainRouter.PathPrefix("/api/warehouses/").Subrouter())
	routers.WheelRouter(mainRouter.PathPrefix("/api/wheels/").Subrouter())
	routers.PropertyRouter(mainRouter.PathPrefix("/api/properties/").Subrouter())
	routers.TransactionRouter(mainRouter.PathPrefix("/api/trx/").Subrouter())
	routers.TransactionDetailRouter(mainRouter.PathPrefix("/api/trx-detail/").Subrouter())
	routers.SaldoRouter(mainRouter.PathPrefix("/api/saldo/").Subrouter())
	routers.ReportRouter(mainRouter.PathPrefix("/api/report/").Subrouter())
	routers.InvoiceRouter(mainRouter.PathPrefix("/api/invoices/").Subrouter())
}

func runServer() {
	cor := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:8081",
			"http://localhost:3000",
			"http://192.168.100.2:3000",
			"http://103.179.56.180",
			"http://ya2yo.com",
			"http://localhost:3000",
			"http://192.168.100.3:3000",
		},

		AllowedMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders: []string{"Accept", "Accept-Language", "Content-Type"},
		//AllowCredentials: true,
		Debug: true,
	})

	handler := cor.Handler(mainRouter)

	fmt.Println("web server run at local: http://localhost:8181/")
	fmt.Println("web server run at: http://pixel.id:8181/")
	log.Fatal(http.ListenAndServe(":8181", handler))
}

func main() {
	loadEnvirontment()
	createRouter()
	loadRouter()
	runServer()
}
