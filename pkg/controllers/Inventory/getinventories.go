package controllers

import (
	// "database/sql"
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

func GetProductsInventory(w http.ResponseWriter, r *http.Request) {
	var products models.Inventory
	var product []models.Inventory
	db := config.Connect()
	defer db.Close()
	rows, err := db.Query("SELECT productname, quantity FROM inventories")
	fmt.Println(rows)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		err = rows.Scan(&products.Product, &products.Quantity)

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
}
