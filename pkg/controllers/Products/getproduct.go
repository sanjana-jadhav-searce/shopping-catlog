package controllers

import (
	"encoding/json"

	"github.com/sanjana-jadhav-searce/shopping-catlog/pkg/config"

	// "io/ioutil"
	// "example.com/pkg/utils"

	"net/http"

	// "strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/sanjana-jadhav-searce/shopping-catlog/pkg/models"
)

func GetProductByName(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()
	var product models.Product
	params := mux.Vars(r)
	result, err := db.Query("SELECT id, name, specification, SKU, category, price FROM products WHERE name=?", params["name"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		err := result.Scan(&product.ID, &product.Name, &product.Specification, &product.SKU, &product.Category, &product.Price)
		if err != nil {
			panic(err.Error())
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
	if product.Name == "" {
		json.NewEncoder(w).Encode(map[string]string{"message": "Product not found"})
	}
}
