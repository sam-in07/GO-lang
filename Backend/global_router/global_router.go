package global_router

import "net/http"

// routes => method => options => cors => sending status
func GlobarRouter(mux *http.ServeMux) http.Handler {
	handelAllReq := func(w http.ResponseWriter, r *http.Request) {
		//halde cors
		w.Header().Set("Access-Control-Allow-Origin", "*") //je access chaibe tare allow kore dibo
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")
		//req methods
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		}
		//options req ase nai Cors handle holo then it goes mux.server
		mux.ServeHTTP(w, r)

	}
	handleReq := http.HandlerFunc(handelAllReq)
	return handleReq
}

//post ..get => mux vitore => mux ke globalrouter a pathie dilam
//preFlight req handle
