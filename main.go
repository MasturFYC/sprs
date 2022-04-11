package main

import (
	"database/sql"
	"os"
	"time"

	"log"

	conn "fyc.com/sprs/controller"
	"fyc.com/sprs/controller/properties"
	"fyc.com/sprs/controller/reports"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func loadEnvirontment() {
	err := godotenv.Load("/home/mastur/.env")
	//err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func database(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}

func runServer() {

	db := InitDatabase()
	defer db().Close()

	port := os.Getenv("PORT")

	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.Default()
	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "POST", "DELETE", "GET"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))
	router.Use(database(db()))

	apiRouter := router.Group("/api")
	{
		authRouter := apiRouter.Group("/auth")
		{
			authRouter.POST("/signup", conn.SignUp)
			authRouter.POST("/signin", conn.SignIn)
		}
		accGroupRouter := apiRouter.Group("/acc-group")
		{
			accGroupRouter.Use(conn.IsAuthorized())
			accGroupRouter.POST("", conn.AccGroupCreate)
			accGroupRouter.PUT("/:id", conn.AccGroupUpdate)
			accGroupRouter.DELETE("/:id", conn.AccGroupDelete)
			accGroupRouter.GET("", conn.AccGroupGetAll)
			accGroupRouter.GET("/types/:id", conn.AccGroupGetTypes)
			accGroupRouter.GET("/all-accounts", conn.AccGroupGetAllAccount)
		}
		accTypeRouter := apiRouter.Group("/acc-type")
		{
			accTypeRouter.POST("", conn.AccTypeCreate)
			accTypeRouter.PUT("/:id", conn.AccTypeUpdate)
			accTypeRouter.DELETE("/:id", conn.AccTypeDelete)
			accTypeRouter.GET("", conn.AccTypeGetAll)
		}
		accCodeRouter := apiRouter.Group("/acc-code")
		{
			accCodeRouter.GET("", conn.AccCodeGetAll)
			accCodeRouter.GET("/search-name/:txt", conn.AccCodeSearchByName)
			accCodeRouter.GET("/group-type/:id", conn.AccCodeGetByType)
			accCodeRouter.GET("/props", conn.AccCodeGetProps)
			accCodeRouter.GET("/item/:id", conn.AccCodeGetItem)
			accCodeRouter.GET("/spec/:id", conn.AccCodeGetSpec)
			accCodeRouter.DELETE("/:id", conn.AccCodeDelete)
			accCodeRouter.POST("", conn.AccCodeCreate)
			accCodeRouter.PUT("/:id", conn.AccCodeUpdate)
		}
		actionRouter := apiRouter.Group("/action")
		{
			actionRouter.POST("", conn.ActionCreate)
			actionRouter.POST("/upload-file/:id", conn.ActionUploadFile)
			actionRouter.PUT("/:id", conn.ActionUpdate)
			actionRouter.DELETE("/:id", conn.ActionDelete)
			actionRouter.GET("/order/:id", conn.ActionGetByOrder) // need attention
			actionRouter.GET("/file/:txt", conn.ActionGetFile)
			actionRouter.GET("/preview/:txt", conn.ActionGetPreview)
		}
		branchRouter := apiRouter.Group("/branch")
		{
			branchRouter.GET("", conn.BranchGetAll)
			branchRouter.GET("/:id", conn.BranchGetItem)
			branchRouter.DELETE("/:id", conn.BranchDelete)
			branchRouter.POST("", conn.BranchCreate)
			branchRouter.PUT("/:id", conn.BranchUpdate)
		}
		customerRouter := apiRouter.Group("/customer")
		{
			customerRouter.POST("", conn.CustomerCreate)
			customerRouter.GET("/:id", conn.CustomerGetItem)
			customerRouter.DELETE("/:id", conn.CustomerDelete)
			customerRouter.PUT("/:id", conn.CustomerUpdate)
		}
		financeGroupRouter := apiRouter.Group("/finance-group")
		{
			financeGroupRouter.GET("", conn.FinanceGroup_GetAll)
			financeGroupRouter.GET("/finances", conn.FinanceGroup_GetFinances)
			financeGroupRouter.GET("/item/:id", conn.FinanceGroup_GetItem) // need attention
			financeGroupRouter.DELETE("/:id", conn.FinanceGroup_Delete)
			financeGroupRouter.POST("", conn.FinanceGroup_Create)
			financeGroupRouter.PUT("/:id", conn.FinanceGroup_Update)
		}
		financeRouter := apiRouter.Group("/finance") // need attention
		{
			financeRouter.GET("", conn.FinanceGetAll)
			financeRouter.GET("/:id", conn.FinanceGetItem)
			financeRouter.DELETE("/:id", conn.FinanceDelete)
			financeRouter.POST("", conn.FinanceCreate)
			financeRouter.PUT("/:id", conn.FinanceUpdate)
		}
		homeAddressRouter := apiRouter.Group("/home-address")
		{
			homeAddressRouter.POST("", conn.CreateHomeAddress)
			homeAddressRouter.GET("/:id", conn.GetHomeAddress)
			homeAddressRouter.PUT("/:id", conn.UpdateHomeAddress)
			homeAddressRouter.DELETE("/:id", conn.DeleteHomeAddress)
		}
		officeAddressRouter := apiRouter.Group("/office-address")
		{
			officeAddressRouter.POST("", conn.CreateOfficeAddress)
			officeAddressRouter.GET("/:id", conn.GetOfficeAddress)
			officeAddressRouter.PUT("/:id", conn.UpdateOfficeAddress)
			officeAddressRouter.DELETE("/:id", conn.DeleteOfficeAddress)
		}
		ktpAddressRouter := apiRouter.Group("/ktp-address")
		{
			ktpAddressRouter.POST("", conn.CreateKTPAddress)
			ktpAddressRouter.GET("/:id", conn.GetKTPAddress)
			ktpAddressRouter.PUT("/:id", conn.UpdateKTPAddress)
			ktpAddressRouter.DELETE("/:id", conn.DeleteKTPAddress)
		}
		postAddressRouter := apiRouter.Group("/post-address")
		{
			postAddressRouter.POST("", conn.CreatePostAddress)
			postAddressRouter.GET("/:id", conn.GetPostAddress)
			postAddressRouter.PUT("/:id", conn.UpdatePostAddress)
			postAddressRouter.DELETE("/:id", conn.DeletePostAddress)
		}
		merkRouter := apiRouter.Group("/merk")
		{
			merkRouter.GET("", conn.MerkGetAll)
			merkRouter.GET("/:id", conn.MerkGetItem)
			merkRouter.DELETE("/:id", conn.MerkDelete)
			merkRouter.POST("", conn.MerkCreate)
			merkRouter.PUT("/:id", conn.MerkUpdate)
		}
		orderRouter := apiRouter.Group("/order")
		{
			orderRouter.GET("", conn.GetOrders)
			orderRouter.POST("/search/:item", conn.SearchOrders)
			orderRouter.GET("/finance/:id", conn.GetOrdersByFinance)
			orderRouter.GET("/branch/:id", conn.GetOrdersByBranch)
			orderRouter.GET("/month/:id", conn.GetOrdersByMonth)
			//orderRouter.("/orders/search/", conn.IsAuthorized(conn.GetOrders))
			orderRouter.GET("/item/:id", conn.GetOrder) // need attention
			orderRouter.DELETE("/:id", conn.DeleteOrder)
			//orderRouter.("/:id/", conn.IsAuthorized(conn.DeleteOrder))
			orderRouter.POST("", conn.CreateOrder)
			orderRouter.PUT("/:id", conn.UpdateOrder)
			orderRouter.GET("/name/seq", conn.Order_GetNameSeq)
			orderRouter.GET("/invoiced/:month/:year/:financeId", conn.Order_GetInvoiced)
		}
		taskRouter := apiRouter.Group("/task")
		{
			taskRouter.POST("", conn.CreateTask)
			taskRouter.GET("/:id", conn.GetTask)
			taskRouter.DELETE("/:id", conn.DeleteTask)
			taskRouter.PUT("/:id", conn.UpdateTask)
		}
		typeRouter := apiRouter.Group("/type")
		{
			typeRouter.GET("", conn.GetTypes)
			typeRouter.GET("/:id", conn.GetType)
			typeRouter.DELETE("/:id", conn.DeleteType)
			typeRouter.POST("", conn.CreateType)
			typeRouter.PUT("/:id", conn.UpdateType)
		}
		unitRouter := apiRouter.Group("/unit")
		{
			unitRouter.GET("/:id", conn.GetUnit)
			unitRouter.DELETE("/:id", conn.DeleteUnit)
			unitRouter.POST("", conn.CreateUnit)
			unitRouter.PUT("/:id", conn.UpdateUnit)
		}
		warehouseRouter := apiRouter.Group("/warehouse")
		{
			warehouseRouter.GET("", conn.GetWarehouses)
			warehouseRouter.POST("", conn.CreateWarehouse)
			warehouseRouter.GET("/:id", conn.GetWarehouse)
			warehouseRouter.DELETE("/:id", conn.DeleteWarehouse)
			warehouseRouter.PUT("/:id", conn.UpdateWarehouse)
		}
		wheelRouter := apiRouter.Group("/wheel")
		{
			wheelRouter.GET("", conn.GetWheels)
			wheelRouter.GET("/:id", conn.GetWheel)
			wheelRouter.DELETE("/:id", conn.DeleteWheel)
			wheelRouter.POST("", conn.CreateWheel)
			wheelRouter.PUT("/:id", conn.UpdateWheel)
		}
		propRouter := apiRouter.Group("/properties")
		{
			propRouter.GET("/product", properties.GetProductsProps)
			propRouter.GET("/category", properties.GetCategoryProps)
		}
		trxRouter := apiRouter.Group("/trx") // need attention
		{
			trxRouter.POST("", conn.TransactionCreate)
			trxRouter.PUT("/:id", conn.TransactionUpdate)
			trxRouter.DELETE("/:id", conn.TransactionDelete)
			trxRouter.GET("", conn.TransactionGetAll)
			trxRouter.POST("/search", conn.TransactionSearch)
			trxRouter.GET("/group/:id", conn.TransactionGetByGroup)
			trxRouter.GET("/month/:id", conn.TransactionGetByMonth)
		}
		trxDetailRouter := apiRouter.Group("/trx-detail")
		{
			trxDetailRouter.GET("/:id", conn.GetTransactionDetails)
			//trxDetailRouter("/props", conn.GetTransactionDetailProps)
			//trxDetailRouter("/:id", conn.GetTransactionDetail)
			trxDetailRouter.DELETE("/:id", conn.DeleteTransactionDetail)
			trxDetailRouter.POST("", conn.CreateTransactionDetail)
			trxDetailRouter.PUT("/:id", conn.UpdateTransactionDetail)
		}
		apiRouter.GET("/saldo", conn.GetRemainSaldo)
		reportRouter := apiRouter.Group("/report")
		{
			reportRouter.GET("/trx/month/:month/:year", conn.GetRepotTrxByMonth)
			reportRouter.GET("/order-status/:finance/:branch/:type/:month/:year/:from/:to", reports.ReportOrder)
			reportRouter.GET("/order-all-waiting/:finance/:branch/:type", reports.ReportOrderAllWaiting)
		}
		invoiceRouter := apiRouter.Group("/invoice") // need attention
		{

			invoiceRouter.POST("/search", conn.Invoice_GetSearch)

			invoiceRouter.GET("/finance/:id", conn.InvoiceGetByFinance)
			invoiceRouter.GET("/month-year/:month/:year", conn.InvoiceGetByMonth) // need attention
			invoiceRouter.GET("/order/:financeId/:id", conn.InvoiceGetOrders)
			invoiceRouter.GET("/download/1/:id", conn.Pdf_GetInvoice)
			//CLIPAN  : 2
			invoiceRouter.GET("/download/2/:id", conn.Clipan_GetInvoice)
			// MTF  : 3
			invoiceRouter.GET("/download/3/:id", conn.Mtf_GetInvoice)

			invoiceRouter.GET("/item/:id", conn.InvoiceGetItem)
			invoiceRouter.GET("", conn.InvoiceGetAll)
			invoiceRouter.POST("", conn.InvoiceCreate)
			invoiceRouter.PUT("/:id", conn.InvoiceUpdate)
			invoiceRouter.DELETE("/:id", conn.InvoiceDelete)
		}
		loanRouter := apiRouter.Group("/loan")
		{
			loanRouter.GET("", conn.LoanGetAll)
			loanRouter.GET("/:id", conn.LoanGetItem)
			loanRouter.DELETE("/:id", conn.LoanDelete)
			loanRouter.POST("", conn.LoanCreate)
			loanRouter.POST("/payment/:id", conn.LoanPayment)
			loanRouter.PUT("/:id", conn.LoanUpdate)
		}
		lentRouter := apiRouter.Group("/lent")
		{
			lentRouter.GET("", conn.LentGetAll)
			lentRouter.GET("/item/:id", conn.LentGetItem) // need attention
			lentRouter.GET("/get/units", conn.LentGetUnits)
			lentRouter.DELETE("/:id", conn.LentDelete)
			lentRouter.POST("", conn.LentCreate)
			lentRouter.POST("/payment/:id", conn.LentPayment)
			lentRouter.PUT("/:id", conn.LentUpdate)
		}
		labaRugiRouter := apiRouter.Group("/labarugi")
		{
			labaRugiRouter.GET("/bydate/:from/:to", conn.LabaRugiGetByDate)
		}
		// testRouter := apiRouter.Group("/test")
		// {
		// 	testRouter.GET("/trx", conn.Test1)
		// }
	}
	router.Run(":" + port) // ":8080"
}

func main() {
	loadEnvirontment()
	//	createRouter()
	runServer()
}

// shraed liberary references
// https://blog.ralch.com/articles/golang-sharing-libraries/
// https://sclem.dev/posts/go-abi/
