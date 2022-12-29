package controllers

import (
	// "database/sql"
	"encoding/json"
	"fmt"

	"github.com/sanjana-jadhav-searce/shopping-catlog/pkg/config"
	"github.com/sanjana-jadhav-searce/shopping-catlog/pkg/models"

	// "io/ioutil"
	// "example.com/pkg/utils"
	"log"
	"net/http"
	// "strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// var db *sql.DB

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	name := r.FormValue("name")

	_, err = db.Exec("INSERT INTO categorymaster(name) VALUES(?)", name)

	if err != nil {
		log.Print(err)
		return
	}
	x := "Inserted Category Data Successfully!"
	json.Marshal(x)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(x)

}

func GetCategoryByName(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()
	var category models.CategoryMaster
	params := mux.Vars(r)
	result, err := db.Query("SELECT name FROM categorymaster WHERE name=?", params["name"])
	fmt.Println(result)
	if err != nil {
		panic(err.Error())
	}
	for result.Next() {
		err := result.Scan(&category.Name)
		if err != nil {
			panic(err.Error())
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(category)
}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	var categories models.CategoryMaster
	var category []models.CategoryMaster
	db := config.Connect()
	defer db.Close()
	rows, err := db.Query("SELECT name FROM categorymaster")
	fmt.Println(rows)
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		err = rows.Scan(&categories.Name)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			category = append(category, categories)
		}
		fmt.Println(category)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(category)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	product := r.FormValue("product")
	rows, err := db.Query("SELECT name FROM categorymaster")
	if product == "" {

		y := "Data Not Found"
		json.Marshal(y)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(y)
		return
	}
	if rows.Next() == false {
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
