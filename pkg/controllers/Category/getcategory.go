package controllers

import (
	// "database/sql"

	"encoding/json"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/sanjana-jadhav-searce/shopping-catlog/pkg/config"
	"github.com/sanjana-jadhav-searce/shopping-catlog/pkg/models"

	// "io/ioutil"
	// "example.com/pkg/utils"
	// "log"
	"net/http"

	// "strconv"

	_ "github.com/go-sql-driver/mysql"
)

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
	if category.Name == "" {
		json.NewEncoder(w).Encode(map[string]string{"message": "Category not found"})
	}
}
