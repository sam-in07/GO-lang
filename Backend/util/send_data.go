package util

import (
	"encoding/json"
	"net/http"
)

func SendData(w http.ResponseWriter, data interface{}, statuscode int) {
	//json theke kono kichu pathiete Encoder
	w.WriteHeader(statuscode)
	encoder := json.NewEncoder(w) //We use json.NewEncoder() to encode and send JSON efficiently to an io.Writer.
	encoder.Encode(data)

}
