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
	"github.com/gorilla/mux"
	"github.com/sanjana-jadhav-searce/shopping-catlog/pkg/models"
)

func GetProductByNameInInventory(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()
	var product models.Inventory
	params := mux.Vars(r)
	result, err := db.Query("SELECT productname, quantity FROM inventories WHERE productname=?", params["name"])
	fmt.Println(result)
	if err != nil {
		panic(err.Error())
	}
	for result.Next() {
		err := result.Scan(&product.Product, &product.Quantity)
		if err != nil {
			panic(err.Error())
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
	if product.Product == "" {
		json.NewEncoder(w).Encode(map[string]string{"message": "Inventory Product not found"})
	}
}
