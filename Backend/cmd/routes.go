package cmd

import (
	"ecomm/hadlers"
	"ecomm/middleware"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /rah", manager.With(
		http.HandlerFunc(hadlers.Test),
	))
	mux.Handle("GET /route", manager.With(
		http.HandlerFunc(hadlers.Test),
	))

	mux.Handle("GET /products", manager.With(
		http.HandlerFunc(hadlers.GetProducts),
	))
	mux.Handle("POST /create-products", manager.With(
		http.HandlerFunc(hadlers.CreateProduct),
	))

	mux.Handle("GET /products/{id}", manager.With(
		http.HandlerFunc(hadlers.GetProductByID),
	))
}
