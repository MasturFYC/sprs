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

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	cor := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:8081",
			"http://localhost:8181",
		},
		AllowedMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders: []string{"Accept", "Accept-Language", "Content-Type"},
		//AllowCredentials: true,
		Debug: true,
	})

	mainRouter := mux.NewRouter()

	routers.ActionRouter(mainRouter.PathPrefix("/api/categories/").Subrouter())
	routers.BranchRouter(mainRouter.PathPrefix("/api/branchs/").Subrouter())
	routers.CustomerRouter(mainRouter.PathPrefix("/api/customers/").Subrouter())
	routers.FinanceRouter(mainRouter.PathPrefix("/api/finances/").Subrouter())
	routers.HomeAddressRouter(mainRouter.PathPrefix("/api/home-address/").Subrouter())
	routers.PostAddressRouter(mainRouter.PathPrefix("/api/post-address/").Subrouter())
	routers.OfficeAddressRouter(mainRouter.PathPrefix("/api/office-address/").Subrouter())
	routers.KtpAddressRouter(mainRouter.PathPrefix("/api/ktp-address/").Subrouter())
	routers.MerkRouter(mainRouter.PathPrefix("/api/merks/").Subrouter())
	routers.WheelRouter(mainRouter.PathPrefix("/api/wheels/").Subrouter())
	routers.TypeRouter(mainRouter.PathPrefix("/api/types/").Subrouter())
	routers.UnitRouter(mainRouter.PathPrefix("/api/units/").Subrouter())
	routers.PropertyRouter(mainRouter.PathPrefix("/api/properties/").Subrouter())

	handler := cor.Handler(mainRouter)

	fmt.Println("web server run at local: http://localhost:8181/")
	fmt.Println("web server run at: http://pixel.id:8181/")
	log.Fatal(http.ListenAndServe(":8181", handler))
}
