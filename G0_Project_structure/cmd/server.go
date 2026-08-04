package cmd

import (
	"fmt"
	"net/http"

	"example.com/g0_project_structure/config"
)

func Server() {
	cnf := config.GetConfig() 


//	mux := http.NewServeMux()
	
   fmt.Println("Server Run on :8080 ")
}