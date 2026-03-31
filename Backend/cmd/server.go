package cmd

import (
	"ecomm/global_router"
	"ecomm/hadlers"
	"fmt"
	"net/http"
)

func Server() {
	mux := http.NewServeMux() //req router

	// mux.Handle("GET /hellow", http.HandlerFunc(helloHandler))
	// mux.Handle("GET /about", http.HandlerFunc(aboutHNandler))

	// mux.Handle("GET /products", http.HandlerFunc(getProducts))
	mux.Handle("GET /products", (http.HandlerFunc(hadlers.GetProducts)))
	//options naile   front a show korbe na products
	// mux.Handle("OPTIONS /products", http.HandlerFunc(getProducts))
	mux.Handle("POST /create-products", (http.HandlerFunc(hadlers.CreateProduct)))
	//create proe show korte ow options lagbe
	// mux.Handle("OPTIONS /create-products", http.HandlerFunc(createProduct))
    mux.Handle("GET /products/{id}", (http.HandlerFunc(hadlers.GetProductByID)))
	fmt.Println("seever running on:8080")
	globarRouter := global_router.GlobarRouter(mux)

	err := http.ListenAndServe(":8080", globarRouter)

	if err != nil {
		fmt.Println("error stRTING server ", err)
	}

}
