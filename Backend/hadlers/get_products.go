package hadlers

import (
	"ecomm/database"
	"ecomm/util"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	// handleCors(w)

	// if r.Method == "OPTIONS" {
	// 	handlePreflightReq(w, r)
	// 	return
	// }

	// if r.Method != "GET" {
	// 	http.Error(
	// 		w, "Plsz give me Get req", 400)
	// 	return
	// }
	//if r.Method != http.MethodGet or if r.Method != "GET"
	// encoder := json.NewEncoder(w)
	//We use json.NewEncoder() to encode and send JSON efficiently to an io.Writer.
	// encoder.Encode(productList)
	//Convert productList to JSON and write it directly to the response.

	util.SendData(w, database.ProductList, 200)
}

//package main kaj korbe na Cuz amr file hadlers vitore tao folder name a hobe
