package config

import (
	"database/sql"
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *sql.DB
)

func Connect() *sql.DB {
	d, err := sql.Open("mysql", "root:Sanju2001@/shoppingcart?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database successfully!")
	db = d
	return db
}

func GetDB() *sql.DB {
	return db
}
