package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter , r *http.Request){
   fmt.Fprint(w,"Hellow World")
}

/*
w http.ResponseWriter → used to send response back to browser
r *http.Request → contains request data (URL, headers, etc.)
*/


func main()  {
	mux := http.NewServeMux() //req router 

	mux.HandleFunc("/hellow",helloHandler)
}  

/*
mux → waiter (routes orders)
/hellow → menu item
helloHandler → chef
fmt.Fprint → serving food

*/