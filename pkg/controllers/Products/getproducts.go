package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/sanjana-jadhav-searce/shopping-catlog/pkg/config"

	// "io/ioutil"
	// "example.com/pkg/utils"

	"net/http"

	// "strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sanjana-jadhav-searce/shopping-catlog/pkg/models"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products models.MinimumProduct
	var product []models.MinimumProduct
	db := config.Connect()
	defer db.Close()
	rows, err := db.Query("SELECT id,name,price,category FROM products limit 20")
	fmt.Println(rows)
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		err = rows.Scan(&products.ID, &products.Name, &products.Price, &products.Category)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			product = append(product, products)
		}
		fmt.Println(product)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(product)
	if products.ID == 0 {
		json.NewEncoder(w).Encode(map[string]string{"message": "Product List is Empty!... Add some Products!!"})
	}
}
