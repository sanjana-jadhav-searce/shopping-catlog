package controllers

import (
	// "database/sql"
	"encoding/json"
	"fmt"

	"github.com/sanjana-jadhav-searce/shopping-catlog/pkg/config"

	// "io/ioutil"
	// "example.com/pkg/utils"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sanjana-jadhav-searce/shopping-catlog/pkg/models"
	// "github.com/gorilla/mux"
)

func CreateCart(w http.ResponseWriter, r *http.Request) {

	db := config.Connect()
	defer db.Close()
	var products models.Cart
	// var cart []models.Cart
	var cartitems = []models.Item{}
	// var addition models.Total
	// var totalcartvalue int64
	fmt.Println(cartitems)
	product := r.FormValue("product")
	quantity := r.FormValue("quantity")
	rows, err := db.Query("SELECT productname, quantity FROM inventories WHERE productname=?", product)
	fmt.Println(rows)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		err = rows.Scan(&products.Product, &products.Quantity)
		cartquantity, err1 := strconv.Atoi(quantity)
		if err1 != nil {
			fmt.Println("error while parsing")
		}
		if err != nil {
			fmt.Println(err.Error())
		} else {
			if products.Quantity < 0 {
				x := map[string]string{"message": "Quantity cannot be Negative"}
				json.NewEncoder(w).Encode(x)
			}
			if products.Quantity >= cartquantity {
				_, err2 := db.Exec("INSERT INTO carts(product, quantity) VALUES(?, ?)", product, quantity)

				if err2 != nil {
					log.Print(err2)
					return
				}
				insertion := "Data Inserted Successfully!"
				json.Marshal(insertion)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(insertion)

			} else {
				x := "Data not inserted"
				json.Marshal(x)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(x)
			}
		}
	}

}
