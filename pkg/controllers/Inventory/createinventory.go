package controllers

import (
	// "database/sql"
	"encoding/json"

	"github.com/sanjana-jadhav-searce/shopping-catlog/pkg/config"

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
