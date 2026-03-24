package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "Hellow World")
}

/*
Sets response header as "application/json" (though you're sending plain text)
Sends "Hellow World" as response
*/

func aboutHNandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hellow I;m asasdjknsd")
}

/*
w http.ResponseWriter → used to send response back to browser
r *http.Request → contains request data (URL, headers, etc.)
*/

type Product struct {
	ID          int   `json:"id"` ///avbe likhle choto hater hoye likha jabe tokhon
	Title       string
	Description string
	Price       float64
	ImgUrl      string
}

//create slice
var productList []Product


func getProducts(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin","*") //je access chaibe tare allow kore dibo
  w.Header().Set("Content-Type","application/json")
	if r.Method != "GET" {
		http.Error(
			w, "Plsz give me Get req", 400)
		return
	}
	//if r.Method != http.MethodGet or if r.Method != "GET"
    encoder := json.NewEncoder(w)     //We use json.NewEncoder() to encode and send JSON efficiently to an io.Writer.
    encoder.Encode(productList)
	//Convert productList to JSON and write it directly to the response.
} 

func main() {
	mux := http.NewServeMux() //req router

	mux.HandleFunc("/", helloHandler)
	mux.HandleFunc("/about", aboutHNandler)
	mux.HandleFunc("/products", getProducts)

	fmt.Println("seever running on:8080")

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		fmt.Println("error stRTING server ", err)
	}

}

/*
mux → waiter (routes orders)
/hellow → menu item
helloHandler → chef
fmt.Fprint → serving food

*/

func init()  {
	prd1 := Product{
		ID: 1,
		Title: "orange",
		Description: "oragwelfhsfdhu",
		Price: 100,
		ImgUrl: "https://www.health.com/thmb/OZgW2YQtFb9qJ3PbySNei3YdgPw=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/Health-Stocksy_txp5e95690asrw300_Medium_934585-e870449543284eed8aa4be52fc09a4ed.jpg",
	}
	prd2 := Product{
		ID: 2,
		Title: "Apple",
		Description: "oragwelfhsfdhu",
		Price: 100,
		ImgUrl: "https://img.freepik.com/free-psd/close-up-delicious-apple_23-2151868338.jpg?semt=ais_hybrid&w=740&q=80",
	}
	prd3 := Product{
		ID: 3,
		Title: "Apple",
		Description: "oragwelfhsfdhu",
		Price: 100,
		ImgUrl: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQZnmH5sKjyuzA7e5fEIWDelduIpJDANHApkZPEjoEnnanvi6qFUn1Ljn_7rD9j6bgbsYMm5fCm_vgvmd7vYfeLezYbl1B5C9-Uro3RY-Y&s=10",
	}

	prd4 := Product{
		ID: 4,
		Title: "Grapes",
		Description: "oragwelfhsfdhu",
		Price: 100,
		ImgUrl: "https://www.foodrepublic.com/img/gallery/15-types-of-grapes-to-know-eat-and-drink/intro-1743188162.jpg",
	}

	prd5 := Product{
		ID: 5,
		Title: "Guava",
		Description: "oragwelfhsfdhu",
		Price: 100,
		ImgUrl: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSLyFmvaUU1q1WUUwQRyoNyCcXiFnahWm_smw&s",
	}

	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)
	productList = append(productList, prd4)
	productList = append(productList, prd5)
}


//Because Go data (like structs, slices) cannot be understood directly outside your program.