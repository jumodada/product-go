package data_models

type Product struct {
	ID           int64  `json: "id" sql:"ID" product:"id"`
	ProductName  string `json: "ProductName" sql:"productName" product:"ProductName"`
	ProductNum   int64  `json: "ProductNum" sql:"productNum" product:"ProductNum"`
	ProductImage string `json: "ProductImage" sql:"productImage" product:"ProductImage"`
	ProductUrl   string `json: "ProductUrl" sql:"productUrl" product:"ProductUrl"`
}
