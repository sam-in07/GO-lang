package hadlers

import (
	"ecomm/database"
	"ecomm/util"
	"net/http"
	"strconv"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pId, err := strconv.Atoi(productID)

	if err != nil {
		http.Error(w, "please give me vadil product id", 400)
		return
	}
	for _, product := range database.ProductList {
		// log.Println(idx,product)
		// log.Println("product_d",id)
		if product.ID == pId {
			util.SendData(w, product, 200)
			return
		}
	}

	util.SendData(w, "Data  pai nai", 404)

}
