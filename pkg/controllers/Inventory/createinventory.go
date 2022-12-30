package controllers

import (
	// "database/sql"
	"encoding/json"

	"github.com/sanjana-jadhav-searce/shopping-catlog/pkg/config"
	"github.com/sanjana-jadhav-searce/shopping-catlog/pkg/models"

	// "io/ioutil"
	// "example.com/pkg/utils"
	"log"
	"net/http"

	// "strconv"

	_ "github.com/go-sql-driver/mysql"
)

func CreateInventoryProduct(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	productname := r.FormValue("productname")
	quantity := r.FormValue("quantity")
	var product models.Inventory

	if product.Product == "" {
		x := "Data Cannot be Created"
		json.NewEncoder(w).Encode(map[string]string{"message": x})

	} else {
		result, err := db.Query("INSERT INTO inventories(productname, quantity) VALUES(?, ?)", productname, quantity)

		if err != nil {
			log.Print(err)
			return
		}

		if result.Next() {
			err := result.Scan(&product.Product)
			if err != nil {
				panic(err.Error())
			}
		}

		x := "Created data successfully into the Products Table"
		json.NewEncoder(w).Encode(map[string]string{"message": x})

	}

}
