package cmd

import (
	"ecomm/middleware"
	"fmt"
	"net/http"
)

func Server() {

	manager := middleware.NewManager()

	manager.Use(middleware.Logger,
		middleware.Hudai)

	mux := http.NewServeMux() //req router
	initRoutes(mux, manager)
	fmt.Println("seever running on:8080")
	globarRouter := middleware.CorsWithPreflight(mux)

	err := http.ListenAndServe(":8080", globarRouter)

	if err != nil {
		fmt.Println("error stRTING server ", err)
	}

}
