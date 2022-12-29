package controllers

import (
	"encoding/json"

	"github.com/sanjana-jadhav-searce/shopping-catlog/pkg/config"

	// "io/ioutil"
	// "example.com/pkg/utils"
	"log"
	"net/http"

	// "strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sanjana-jadhav-searce/shopping-catlog/pkg/models"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()
	var demo models.Product
	product := r.FormValue("product")
	if product == "" {

		y := "Please Enter Required Feilds"
		json.NewEncoder(w).Encode(map[string]string{"message": y})
		return
	}
	rows, err := db.Query("SELECT id, name, specification, SKU, category, price FROM products WHERE name=?", product)
	if err != nil {
		log.Print(err)
	}
	if !rows.Next() {
		rows.Scan(&demo.Name)
		z := "Invalid Inventory Product Reference"
		json.NewEncoder(w).Encode(map[string]string{"message": z})
		return

	}
	result, err := db.Exec("DELETE FROM products WHERE name=?", product)
	if err != nil {
		log.Print(err)
		return
	}
	rows_affected, err := result.RowsAffected()
	if err != nil {
		log.Print(err)
		return
	}
	if rows_affected != 0 {
		if err != nil {
			log.Print(err)
			return
		} else {
			msg := "Product item deleted successfully"
			json.NewEncoder(w).Encode(map[string]string{"message": msg})
			return

		}
	} else {
		msg := "Product is not found!"
		json.NewEncoder(w).Encode(map[string]string{"message": msg})
		return

	}
}
