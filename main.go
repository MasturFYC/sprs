package main

import (
	"time"

	"log"

	conn "fyc.com/sprs/controller"
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
			accTypeRouter.DELETE("", conn.AccTypeDelete)
			accTypeRouter.GET("", conn.AccTypeGetAll)
		}
		accCodeRouter := apiRouter.Group("/acc-code")
		{
			accCodeRouter.POST("", conn.AccCodeCreate)
			accCodeRouter.PUT("/:id", conn.AccCodeUpdate)
			accCodeRouter.DELETE("/:id", conn.AccCodeDelete)
			accCodeRouter.GET("", conn.AccCodeGetAll)
			accCodeRouter.GET("/search-name/:txt", conn.AccCodeSearchByName)
			accCodeRouter.GET("/group-type/:id", conn.AccCodeGetByType)
			accCodeRouter.GET("/props", conn.AccCodeGetProps)
			accCodeRouter.GET("/spec/:id", conn.AccCodeGetSpec)
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
		invoiceRouter := apiRouter.Group("/invoice") // need attention
		{

			invoiceRouter.GET("/finance/:id", conn.Invoice_GetByFinance)
			invoiceRouter.GET("/month-year/:month/:year", conn.Invoice_GetByMonth) // need attention

			invoiceRouter.GET("/order/:financeId/:id", conn.Invoice_GetOrders)
			invoiceRouter.POST("/search", conn.Invoice_GetSearch)

			invoiceRouter.GET("/download/1/:id", conn.Pdf_GetInvoice)

			//CLIPAN  : 2
			invoiceRouter.GET("/download/2/:id", conn.Clipan_GetInvoice)
			// MTF  : 3
			invoiceRouter.GET("/download/3/:id", conn.Mtf_GetInvoice)

			invoiceRouter.GET("/item/:id", conn.Invoice_GetItem)
			invoiceRouter.GET("", conn.Invoice_GetAll)
			invoiceRouter.POST("", conn.Invoice_Create)
			invoiceRouter.PUT("/:id", conn.Invoice_Update)
			invoiceRouter.DELETE("/:id", conn.Invoice_Delete)
		}
		financeRouter := apiRouter.Group("/finance") // need attention
		{
			financeRouter.GET("", conn.FinanceGetAll)
			financeRouter.GET("/:id", conn.FinanceGetItem)
			financeRouter.DELETE("/:id", conn.FinanceDelete)
			financeRouter.POST("", conn.FinanceCreate)
			financeRouter.PUT("/:id", conn.FinanceUpdate)
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
			lentRouter.DELETE("/:id", conn.LentDelete)
			lentRouter.POST("", conn.LentCreate)
			lentRouter.PUT("/:id", conn.LentUpdate)
			lentRouter.GET("/item/:id", conn.LentGetItem) // need attention
			lentRouter.POST("/payment/:id", conn.LentPayment)
			lentRouter.GET("/get/units", conn.LentGetUnits)
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
	}

	router.Run() // ":8080"

}

func main() {
	loadEnvirontment()
	//	createRouter()
	runServer()
}
