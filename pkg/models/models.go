package models

import (
	"database/sql"

	"github.com/jinzhu/gorm"
	"github.com/sanjana-jadhav-searce/shopping-catlog/pkg/config"
)

var db *sql.DB

type Product struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Specification string `json:"specification"`
	SKU           string `json:"SKU"`
	Category      string `json:"category"`
	Price         int64  `json:"price"`
}
type MinimumProduct struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int64  `json:"price"`
}

type CategoryMaster struct {
	Name string `json:"categoryname"`
}

type Inventory struct {
	Product  string `json:"productname"`
	Quantity int    `json:"quantity"`
}

type Cart struct {
	Product  string `json:"product"`
	Quantity int    `json:"quantity"`
}

type Item struct {
	C []Cart `json:"item"`
}
type DemoCart struct {
	Product  string `json:"product"`
	Quantity int64  `json:"quantity"`
}
type Demo struct {
	D []DemoCart `json:"demoitem"`
}
type Total struct {
	Product  string `json:"product"`
	Quantity int64  `json:"quantity"`
	Price    int64  `json:"price"`
}

var db1 *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
	// db.AutoMigrate(&Product{})
}
func (b *Cart) CreateCart() *Cart {

	db1.NewRecord(b)
	db1.Create(&b)
	return b
}
