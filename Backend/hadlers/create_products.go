package hadlers

import (
	"ecomm/database"
	
	"encoding/json"
	"fmt"
	"net/http"

	//go.mod theke "x" ke copy kore "x/foldername"
	"ecomm/util"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	// handleCors(w)
	// if r.Method == "OPTIONS" {
	// 	handlePreflightReq(w, r)
	// 	return
	// }

	//apply advance routing this dont need
	// if r.Method != "POST" {
	// 	http.Error(w, "PLZ! Give me post req", 400)
	// }

	//r.body => desc , imageUrl , price , title => Product ar akta instance => pOrduct list => append
	/*
		    1.take body information (desciption , imgurl, proice , title) from r.body
			2.create an instance using Product Struct with body information
			3.append the  instance into product list

	*/

	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Pls give me valid JSon", 400)
		return
	}
	/*
		Reads the request body (JSON)
		Converts it into a Product struct (newProduct)
		Read the JSON data from the request body and convert it into a Go Product struct,
		storing the result in newProduct.
	*/
	newProduct.ID = len(database.ProductList) + 1

	database.ProductList = append(database.ProductList, newProduct)

	util.SendData(w, newProduct, 201)

}
