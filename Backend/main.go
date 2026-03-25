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
	ID          int `json:"id"` ///avbe likhle choto hater hoye likha jabe tokhon
	Title       string
	Description string
	Price       float64
	ImgUrl      string
}

// create slice
var productList []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	handleCors(w)

	handlePreflightReq(w, r)

	if r.Method != "GET" {
		http.Error(
			w, "Plsz give me Get req", 400)
		return
	}
	//if r.Method != http.MethodGet or if r.Method != "GET"
	// encoder := json.NewEncoder(w)
	//We use json.NewEncoder() to encode and send JSON efficiently to an io.Writer.
	// encoder.Encode(productList)
	//Convert productList to JSON and write it directly to the response.

	sendData(w, productList, 200)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	handlePreflightReq(w, r)
	if r.Method != "POST" {
		http.Error(w, "PLZ! Give me post req", 400)
	}

	//r.body => desc , imageUrl , price , title => Product ar akta instance => pOrduct list => append
	/*
		    1.take body information (desciption , imgurl, proice , title) from r.body
			2.create an instance using Product Struct with body information
			3.append the  instance into product list

	*/

	var newProduct Product

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
	newProduct.ID = len(productList) + 1

	productList = append(productList, newProduct)

	sendData(w, newProduct, 201)

}

func handleCors(w http.ResponseWriter) {
	//allow accesss
	w.Header().Set("Access-Control-Allow-Origin", "*") //je access chaibe tare allow kore dibo
	w.Header().Set("Access-Control-Allow-Methods", "GET , POST , PUT , PATCH , DELETE , OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
}

func handlePreflightReq(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)

	}
}

func sendData(w http.ResponseWriter, data interface{}, statuscode int) {
	//json theke kono kichu pathiete Encoder
	w.WriteHeader(statuscode)
	encoder := json.NewEncoder(w) //We use json.NewEncoder() to encode and send JSON efficiently to an io.Writer.
	encoder.Encode(data)

}

func main() {
	mux := http.NewServeMux() //req router

	mux.HandleFunc("/", helloHandler)
	mux.HandleFunc("/about", aboutHNandler)
	mux.HandleFunc("/products", getProducts)
	mux.HandleFunc("/create-products", createProduct)

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

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "orange",
		Description: "oragwelfhsfdhu",
		Price:       100,
		ImgUrl:      "https://www.health.com/thmb/OZgW2YQtFb9qJ3PbySNei3YdgPw=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/Health-Stocksy_txp5e95690asrw300_Medium_934585-e870449543284eed8aa4be52fc09a4ed.jpg",
	}
	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "oragwelfhsfdhu",
		Price:       100,
		ImgUrl:      "https://img.freepik.com/free-psd/close-up-delicious-apple_23-2151868338.jpg?semt=ais_hybrid&w=740&q=80",
	}
	prd3 := Product{
		ID:          3,
		Title:       "Apple",
		Description: "oragwelfhsfdhu",
		Price:       100,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQZnmH5sKjyuzA7e5fEIWDelduIpJDANHApkZPEjoEnnanvi6qFUn1Ljn_7rD9j6bgbsYMm5fCm_vgvmd7vYfeLezYbl1B5C9-Uro3RY-Y&s=10",
	}

	prd4 := Product{
		ID:          4,
		Title:       "Grapes",
		Description: "oragwelfhsfdhu",
		Price:       100,
		ImgUrl:      "https://www.foodrepublic.com/img/gallery/15-types-of-grapes-to-know-eat-and-drink/intro-1743188162.jpg",
	}

	prd5 := Product{
		ID:          5,
		Title:       "Guava",
		Description: "oragwelfhsfdhu",
		Price:       100,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSLyFmvaUU1q1WUUwQRyoNyCcXiFnahWm_smw&s",
	}

	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)
	productList = append(productList, prd4)
	productList = append(productList, prd5)
}

//Because Go data (like structs, slices) cannot be understood directly outside your program.
