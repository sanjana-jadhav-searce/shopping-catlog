package controllers

import (
	// "database/sql"

	"encoding/json"
	"fmt"

	"github.com/sanjana-jadhav-searce/shopping-catlog/pkg/config"
	"github.com/sanjana-jadhav-searce/shopping-catlog/pkg/models"

	// "io/ioutil"
	// "example.com/pkg/utils"

	"net/http"

	// "strconv"

	_ "github.com/go-sql-driver/mysql"
)

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
