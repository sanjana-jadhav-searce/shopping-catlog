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
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)

	if err != nil {
		panic(err)
	}

	name := r.FormValue("name")
	// specification := r.FormValue("specification")
	// SKU := r.FormValue("SKU")
	// category := r.FormValue("category")
	// price := r.FormValue("price")
	id := r.FormValue("id")

	_, err = db.Exec("UPDATE products SET name=? WHERE id=?", name, id)

	if err != nil {
		log.Print(err)
	}

	x := "Updated data successfully!"
	json.NewEncoder(w).Encode(map[string]string{"message": x})
}
