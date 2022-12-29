package controllers

import (
	// "database/sql"

	"fmt"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sanjana-jadhav-searce/shopping-catlog/pkg/config"
	"github.com/sanjana-jadhav-searce/shopping-catlog/pkg/models"
	// "github.com/gorilla/mux"
)

func MultipleCart(quantity string, product string) map[string]string {

	db := config.Connect()
	defer db.Close()

	var products models.Cart
	var cart []models.Cart
	// var cartitems = []models.Item{}

	if quantity == "" || product == "" {
		return map[string]string{"message": "ref / quantity / product_id missing in the url"}
	}
	rows, err := db.Query("SELECT productname, quantity FROM inventories WHERE productname=?", product)
	if err != nil {
		fmt.Println(err)
	}
	if !rows.Next() {
		return map[string]string{"message": "Invalid cart_reference"}
	}
	for rows.Next() {
		err = rows.Scan(&products.Product, &products.Quantity)
		cartquantity, err1 := strconv.ParseInt(quantity, 0, 0)
		if err1 != nil {
			fmt.Println("error while parsing")
		}
		if err != nil {
			fmt.Println(err.Error())
		} else {
			if products.Quantity >= cartquantity {
				_, err2 := db.Exec("INSERT INTO carts(product, quantity) VALUES(?, ?)", product, quantity)

				if err2 != nil {
					log.Print(err2)
				}
				rows, err3 := db.Query("SELECT product, quantity FROM carts")
				if err3 != nil {
					fmt.Println(err3)
				}
				for rows.Next() {
					err3 = rows.Scan(&products.Product, &products.Quantity)
					if err3 != nil {
						fmt.Println(err3.Error())
					} else {
						cart = append(cart, products)
					}
					fmt.Println(cart)
				}
			} else {
				return map[string]string{"message": "Inventory Quantity is greater than required quantity"}
			}
		}
	}
	return map[string]string{"message": "Item was added to the cart"}

}
