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

	// "strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/sanjana-jadhav-searce/shopping-catlog/pkg/models"
)

// var db *sql.DB

func CreateInventoryProduct(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	productname := r.FormValue("productname")
	quantity := r.FormValue("quantity")

	_, err = db.Exec("INSERT INTO inventories(productname, quantity) VALUES(?, ?)", productname, quantity)

	if err != nil {
		log.Print(err)
		return
	}
	x := "Inserted product into inventories successfully!"
	json.Marshal(x)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(x)

}

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
}

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

func UpdateProductInventory(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)

	if err != nil {
		panic(err)
	}

	productname := r.FormValue("productname")
	quantity := r.FormValue("quantity")

	_, err = db.Exec("UPDATE inventories SET quantity=? WHERE productname=?", quantity, productname)

	if err != nil {
		log.Print(err)
	}

	x := "Updated Inventory data successfully!"
	json.Marshal(x)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(x)
}

func DeleteProductInventory(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()
	var demo models.Inventory
	product := r.FormValue("product")
	rows, err := db.Query("SELECT productname, quantity FROM inventories WHERE productname=?", product)
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
		rows.Scan(&demo.Product)
		z := "Invalid Inventory Product Reference"
		json.Marshal(z)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(z)
		return

	}
	result, err := db.Exec("DELETE FROM inventories WHERE productname=?", product)
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
			msg := "Inventory item deleted successfully"
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
