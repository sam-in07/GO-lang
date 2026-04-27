package cmd

import (
	"ecomm/global_router"
	"ecomm/hadlers"
	"ecomm/middleware"
	"fmt"
	"net/http"
)

func Server() {

	manager := middleware.NewManager()

	mux := http.NewServeMux() //req router
	mux.Handle("GET /rah", manager.With(
		http.HandlerFunc(hadlers.Test),
		middleware.Logger,
		middleware.Hudai,
	))
	mux.Handle("GET /route", manager.With(
		http.HandlerFunc(hadlers.Test),
		middleware.Logger,
		middleware.Hudai,
	))

	mux.Handle("GET /products", manager.With(
		http.HandlerFunc(hadlers.GetProducts),
		middleware.Logger,
		middleware.Hudai,
	))
	mux.Handle("POST /create-products", manager.With(
		http.HandlerFunc(hadlers.CreateProduct),
		middleware.Logger,
		middleware.Hudai,
	))

	mux.Handle("GET /products/{id}", manager.With(
		http.HandlerFunc(hadlers.GetProductByID),
		middleware.Logger,
		middleware.Hudai,
	))
	fmt.Println("seever running on:8080")
	globarRouter := global_router.GlobarRouter(mux)

	err := http.ListenAndServe(":8080", globarRouter)

	if err != nil {
		fmt.Println("error stRTING server ", err)
	}

}
