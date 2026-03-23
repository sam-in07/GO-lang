package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type","application/json")
   fmt.Fprint(w,"Hellow World")
}

/*
Sets response header as "application/json" (though you're sending plain text)
Sends "Hellow World" as response
*/


func aboutHNandler(w http.ResponseWriter , r *http.Request){
   fmt.Fprint(w,"Hellow I;m asasdjknsd")
}

/*
w http.ResponseWriter → used to send response back to browser
r *http.Request → contains request data (URL, headers, etc.)
*/


func main()  {
	mux := http.NewServeMux() //req router 

	mux.HandleFunc("/",helloHandler)
	mux.HandleFunc("/about",aboutHNandler)

	fmt.Println("seever running on:8080")

	err := http.ListenAndServe(":8080",mux)

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