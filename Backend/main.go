package main

import (
	"ecomm/cmd"
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

// func handleCors(w http.ResponseWriter) {
// 	//allow accesss
// 	w.Header().Set("Access-Control-Allow-Origin", "*") //je access chaibe tare allow kore dibo
// 	w.Header().Set("Access-Control-Allow-Methods", "GET , POST , PUT , PATCH , DELETE , OPTIONS")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 	w.Header().Set("Content-Type", "application/json")
// }

// func handlePreflightReq(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "OPTIONS" {
// 		w.WriteHeader(200)

// 	}
// }

func main() {
	cmd.Server()
}

/*
mux → waiter (routes orders)
/hellow → menu item
helloHandler → chef
fmt.Fprint → serving food

*/

// func handleCorsMiddleware(next http.Handler) http.Handler {
// 	handleCors := func(w http.ResponseWriter, r *http.Request) {

// 		next.ServeHTTP(w, r)

// 		//get product handler
// 	}
// 	handle := http.HandlerFunc(handleCors)
// 	return handle
// }

// Because Go data (like structs, slices) cannot be understood directly outside your program.
