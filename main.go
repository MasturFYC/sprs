package main

import (
	"time"

	"log"

	conn "fyc.com/sprs/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

// var (
// 	mainRouter *mux.Router
// )

// func createRouter() {
// 	mainRouter = mux.NewRouter()
// }

func loadEnvirontment() {
	// home, err := os.UserHomeDir()
	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }
	// env_path := filepath.Join(home, ".env.sprs")
	err := godotenv.Load("/home/mastur/.env")
	//err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func loadRouter() {

	// routers.InitializeRoute(mainRouter)

	// routers.AccGroupRouter(mainRouter.PathPrefix("/api/acc-group").Subrouter())
	// routers.AccountTypeRouter(mainRouter.PathPrefix("/api/acc-type").Subrouter())
	// routers.AccountCodeRouter(mainRouter.PathPrefix("/api/acc-code").Subrouter())
	// routers.ActionRouter(mainRouter.PathPrefix("/api/actions/").Subrouter())
	// routers.BranchRouter(mainRouter.PathPrefix("/api/branchs").Subrouter())
	// routers.CustomerRouter(mainRouter.PathPrefix("/api/customers").Subrouter())
	// routers.FinanceGroupRouter(mainRouter.PathPrefix("/api/finance-group").Subrouter())
	// routers.FinanceRouter(mainRouter.PathPrefix("/api/finances/").Subrouter())

	// routers.HomeAddressRouter(mainRouter.PathPrefix("/api/home-address").Subrouter())
	// routers.KtpAddressRouter(mainRouter.PathPrefix("/api/ktp-address").Subrouter())
	// routers.OfficeAddressRouter(mainRouter.PathPrefix("/api/office-address").Subrouter())
	// routers.PostAddressRouter(mainRouter.PathPrefix("/api/post-address").Subrouter())

	// routers.MerkRouter(mainRouter.PathPrefix("/api/merks/").Subrouter())
	// routers.OrderRouter(mainRouter.PathPrefix("/api/orders").Subrouter())

	// // routers.ReceivableRouter(mainRouter.PathPrefix("/api/receivables/").Subrouter())
	// routers.TaskRouter(mainRouter.PathPrefix("/api/tasks").Subrouter())
	// routers.TypeRouter(mainRouter.PathPrefix("/api/types").Subrouter())
	// routers.UnitRouter(mainRouter.PathPrefix("/api/units/").Subrouter())
	// routers.WarehouseRouter(mainRouter.PathPrefix("/api/warehouses/").Subrouter())
	// routers.WheelRouter(mainRouter.PathPrefix("/api/wheels").Subrouter())
	// routers.PropertyRouter(mainRouter.PathPrefix("/api/properties/").Subrouter())
	// routers.TransactionRouter(mainRouter.PathPrefix("/api/trx/").Subrouter())
	// routers.TransactionDetailRouter(mainRouter.PathPrefix("/api/trx-detail/").Subrouter())
	// routers.SaldoRouter(mainRouter.PathPrefix("/api/saldo/").Subrouter())
	// routers.ReportRouter(mainRouter.PathPrefix("/api/report/").Subrouter())
	// routers.InvoiceRouter(mainRouter.PathPrefix("/api/invoices/").Subrouter())

	// routers.LoanRouter(mainRouter.PathPrefix("/api/loans").Subrouter())
	// routers.LentRouter(mainRouter.PathPrefix("/api/lents").Subrouter())

}

func runServer() {

	router := gin.Default()
	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:3000"},
		AllowMethods:  []string{"PUT", "PATCH"},
		AllowHeaders:  []string{"Origin"},
		ExposeHeaders: []string{"Content-Length"},
		// AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

	apiRouter := router.Group("/api")
	{
		accGroupRouter := apiRouter.Group("/acc-group")
		{
			accGroupRouter.GET("", conn.GetAccGroups)
			accGroupRouter.GET("/types/:id", conn.Group_GetTypes)
		}
	}

	router.Run(":8181")

}

func main() {
	loadEnvirontment()
	//	createRouter()
	runServer()
}
