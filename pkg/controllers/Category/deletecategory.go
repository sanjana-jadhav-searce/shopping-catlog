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

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()
	var category models.CategoryMaster
	product := r.FormValue("product")
	rows, err := db.Query("SELECT name FROM categorymaster")
	if err != nil {
		log.Print(err)
	}
	if product == "" {

		y := "Data Not Found"
		json.Marshal(y)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(y)
		return
	}
	if !rows.Next() {
		rows.Scan(&category.Name)
		z := "Invalid Category Reference"
		json.Marshal(z)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(z)
		return

	}
	result, err := db.Exec("DELETE FROM categorymaster WHERE name=?", product)
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
			msg := "Category item deleted successfully"
			json.Marshal(msg)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(msg)
			return

		}
	} else {
		msg := "Product is not found!"
		json.Marshal(msg)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(msg)
		return

	}
}