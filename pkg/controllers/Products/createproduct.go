package controllers

import (
	"encoding/json"
	"log"

	// "io/ioutil"
	// "example.com/pkg/utils"

	"net/http"

	// "strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sanjana-jadhav-searce/shopping-catlog/pkg/config"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	id := r.FormValue("id")
	name := r.FormValue("name")
	specification := r.FormValue("specification")
	SKU := r.FormValue("SKU")
	category := r.FormValue("category")
	price := r.FormValue("price")

	_, err = db.Exec("INSERT INTO products(id, name, specification, SKU, category, price) VALUES(?,?, ?, ?, ?, ?)", id, name, specification, SKU, category, price)

	if err != nil {
		log.Print(err)
		return
	}
	x := "Created data successfully into the Products Table"
	json.NewEncoder(w).Encode(map[string]string{"message": x})

	// requestBody, _ := io.ReadAll(r.Body)
	// var person models.Product
	// json.Unmarshal(requestBody, &person)
	// fmt.Println(person)
	// result, err := db.Query("Insert into products(id, name, specification, sku, category, price) values(?,?,?,?,?)", person.ID, person.Name, person.Specification, person.SKU, person.Category, person.Price)
	// if err != nil {
	// 	log.Print(err)
	// }
	// for result.Next() {
	// 	result.Scan(person.Name, person.Specification, person.SKU, person.Category, person.Price, person.ID)
	// }
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode(person)

}
