package model

type Product struct {
	ID           int64  `json:"id" sql:"id" form:"id" `
	ProductName  string `json:"ProductName" sql:"productName" form:"ProductName"`
	ProductNum   int64  `json:"ProductNum" sql:"productNum" form:"ProductNum"`
	ProductImage string `json:"ProductImage" sql:"productImage" form:"ProductImage"`
	ProductUrl   string `json:"ProductUrl" sql:"productUrl" form:"ProductUrl"`
}
