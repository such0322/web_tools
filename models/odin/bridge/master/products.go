package master

import "time"

type Product struct {
	Model
	ID               int `gorm:"primary_key"`
	Platform         string
	AppId            string
	Store            string
	ProductId        string
	Name             string
	Type             string `gorm:"type:enum('consumable','non-consumable','subscription')"`
	Payment          int
	Coin             int
	SellingType      string `gorm:"DEFAULT:'normal'"`
	PurchaseType     string `gorm:"DEFAULT:'normal'"`
	StartDate        time.Time
	EndDate          time.Time
	Enabled          int
	PurchaseCountMax int
	ReplaceProductId string
}

func (Product) TableName() string {
	return "products"
}

func LoadProduct(id int) *Product {
	var product Product
	DB.First(&product, id)
	return &product
}

func GetProducts() []Product {
	var products []Product
	DB.Find(&products)
	return products
}

type ProductList map[int]Product

func GetProductList() ProductList {
	products := GetProducts()
	productList := ProductList{}
	for _, product := range products {
		productList[product.ID] = product
	}
	return productList
}
