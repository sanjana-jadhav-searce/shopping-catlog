package controllers

import (
	// "database/sql"
	"encoding/json"
	"fmt"

	"github.com/sanjana-jadhav-searce/shopping-catlog/pkg/config"

	// "io/ioutil"
	// "example.com/pkg/utils"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sanjana-jadhav-searce/shopping-catlog/pkg/models"
	// "github.com/gorilla/mux"
)

func GetCart(w http.ResponseWriter, r *http.Request) {

	db := config.Connect()
	defer db.Close()
	var products models.Cart
	var cart []models.Cart

	w.Header().Set("Content-Type", "application/json")
	message := "All Cart Items ---> "
	json.Marshal(message)
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
	json.NewEncoder(w).Encode(message)
	json.NewEncoder(w).Encode(cart)
}
