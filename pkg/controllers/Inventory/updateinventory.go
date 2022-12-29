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
