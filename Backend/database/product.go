package database

// create slice
var ProductList []Product


//productList prive hoi gesa pblic ProductList


type Product struct {
	ID          int `json:"id"` ///avbe likhle choto hater hoye likha jabe tokhon
	Title       string
	Description string
	Price       float64
	ImgUrl      string
}




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

	ProductList = append(ProductList, prd1)
	ProductList = append(ProductList, prd2)
	ProductList = append(ProductList, prd3)
	ProductList = append(ProductList, prd4)
	ProductList = append(ProductList, prd5)
}