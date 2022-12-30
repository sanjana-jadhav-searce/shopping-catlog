package controllers

import (

	// "io/ioutil"
	// "example.com/pkg/utils"

	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	// "strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sanjana-jadhav-searce/shopping-catlog/pkg/config"
	"github.com/sanjana-jadhav-searce/shopping-catlog/pkg/models"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	// db := config.Connect()
	// defer db.Close()

	// err := r.ParseMultipartForm(4096)
	// if err != nil {
	// 	panic(err)
	// }

	// id := r.FormValue("id")
	// name := r.FormValue("name")
	// specification := r.FormValue("specification")
	// SKU := r.FormValue("SKU")
	// category := r.FormValue("category")
	// price := r.FormValue("price")

	// _, err = db.Exec("INSERT INTO products(id, name, specification, SKU, category, price) VALUES(?,?, ?, ?, ?, ?)", id, name, specification, SKU, category, price)

	// if err != nil {
	// 	log.Print(err)
	// 	return
	// }
	// x := "Created data successfully into the Products Table"
	// json.NewEncoder(w).Encode(map[string]string{"message": x})

	requestBody, _ := io.ReadAll(r.Body)
	var product models.Product
	json.Unmarshal(requestBody, &product)
	fmt.Println(product)
	result, err := config.GetDB().Query("INSERT INTO products(id, name, specification, sku, category, price) VALUES(?, ?, ?, ?, ?)", product.ID, product.Name, product.Specification, product.SKU, product.Category, product.Price)
	if err != nil {
		log.Print(err)
	}
	for result.Next() {
		result.Scan(product.ID, product.Name, product.Specification, product.SKU, product.Category, product.Price)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)

}
