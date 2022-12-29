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
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sanjana-jadhav-searce/shopping-catlog/pkg/models"
	// "github.com/gorilla/mux"
)

// var db *sql.DB

func AddItemToCart(w http.ResponseWriter, r *http.Request) {

	db := config.Connect()
	defer db.Close()
	var products models.Cart
	var cart []models.Cart
	var cartitems = []models.Item{}
	var addition models.Total
	var totalcartvalue int64
	fmt.Println(cartitems)
	product := r.FormValue("product")
	quantity := r.FormValue("quantity")
	rows, err := db.Query("SELECT productname, quantity FROM inventories WHERE productname=?", product)
	fmt.Println(rows)
	if err != nil {
		fmt.Println(err)
	}
	if !rows.Next() {
		insertion := "Product Not available in the Inventory!"
		json.Marshal(insertion)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(insertion)
	}
	for rows.Next() {
		err = rows.Scan(&products.Product, &products.Quantity)
		cartquantity, err1 := strconv.ParseInt(quantity, 0, 0)
		if err1 != nil {
			fmt.Println("error while parsing")
		}
		if err != nil {
			fmt.Println(err.Error())
		} else {
			if products.Quantity >= cartquantity {
				_, err2 := db.Exec("INSERT INTO carts(product, quantity) VALUES(?, ?)", product, quantity)

				if err2 != nil {
					log.Print(err2)
					return
				}
				insertion := "Data Inserted Successfully!"
				json.Marshal(insertion)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(insertion)
				message := "All Cart Items ---> "
				json.Marshal(message)
				rows, err3 := db.Query("SELECT product, quantity FROM carts")
				if err3 != nil {
					fmt.Println(err3)
				}
				for rows.Next() {
					err3 = rows.Scan(&products.Product, &products.Quantity)
					if err3 != nil {
						fmt.Println(err3.Error())
					} else {
						cart = append(cart, products)
					}
					fmt.Println(cart)
				}
				rows, err4 := db.Query("SELECT products.name,carts.quantity,products.price FROM products Cross JOIN carts ON products.name = carts.product ORDER BY products.name;")
				fmt.Println(rows)
				if err4 != nil {
					fmt.Println(err4)
				}

				for rows.Next() {
					err4 = rows.Scan(&addition.Product, &addition.Quantity, &addition.Price)
					totalcartvalue += (addition.Quantity * addition.Price)
					if err4 != nil {
						log.Print(err4)
					}
				}
				json.Marshal(totalcartvalue)
				message2 := "Total Cart Value ---> "
				json.NewEncoder(w).Encode(message)
				json.NewEncoder(w).Encode(cart)
				json.NewEncoder(w).Encode(message2)
				json.NewEncoder(w).Encode(totalcartvalue)
			} else {
				x := "Data not inserted"
				json.Marshal(x)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(x)
			}
		}
	}

}

func CreateCart(w http.ResponseWriter, r *http.Request) {

	db := config.Connect()
	defer db.Close()
	var products models.Cart
	// var cart []models.Cart
	var cartitems = []models.Item{}
	// var addition models.Total
	// var totalcartvalue int64
	fmt.Println(cartitems)
	product := r.FormValue("product")
	quantity := r.FormValue("quantity")
	rows, err := db.Query("SELECT productname, quantity FROM inventories WHERE productname=?", product)
	fmt.Println(rows)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		err = rows.Scan(&products.Product, &products.Quantity)
		cartquantity, err1 := strconv.ParseInt(quantity, 0, 0)
		if err1 != nil {
			fmt.Println("error while parsing")
		}
		if err != nil {
			fmt.Println(err.Error())
		} else {
			if products.Quantity >= cartquantity {
				_, err2 := db.Exec("INSERT INTO carts(product, quantity) VALUES(?, ?)", product, quantity)

				if err2 != nil {
					log.Print(err2)
					return
				}
				insertion := "Data Inserted Successfully!"
				json.Marshal(insertion)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(insertion)

			} else {
				x := "Data not inserted"
				json.Marshal(x)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(x)
			}
		}
	}

}

func GetCart(w http.ResponseWriter, r *http.Request) {

	db := config.Connect()
	defer db.Close()
	var products models.Cart
	var cart []models.Cart

	w.Header().Set("Content-Type", "application/json")
	message := "All Cart Items ---> "
	json.Marshal(message)
	rows, err3 := db.Query("SELECT product, quantity FROM carts")
	if err3 != nil {
		fmt.Println(err3)
	}
	for rows.Next() {
		err3 = rows.Scan(&products.Product, &products.Quantity)
		if err3 != nil {
			fmt.Println(err3.Error())
		} else {
			cart = append(cart, products)
		}
		fmt.Println(cart)
	}
	json.NewEncoder(w).Encode(message)
	json.NewEncoder(w).Encode(cart)
}

func DeleteCart(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	product := r.FormValue("product")
	rows, err := db.Query("SELECT product, quantity FROM carts WHERE product=?", product)
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
		z := "Invalid Cart Reference"
		json.Marshal(z)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(z)
		return

	}
	result, err := db.Exec("DELETE FROM carts WHERE product=?", product)
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
			msg := "Cart item deleted successfully"
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

func AddItemsToCart(w http.ResponseWriter, r *http.Request) {
	response := []map[string]any{}
	request_body := []map[string]int{}

	err := json.NewDecoder(r.Body).Decode(&request_body)
	if err != nil {
		return
	}

	for _, v := range request_body {
		new_response_item := map[string]any{}
		product := v["product"]
		quantity := v["quantity"]
		new_response_item["product"] = product
		new_response_item["quantity"] = quantity

		new_response_item["message"] = MultipleCart(fmt.Sprint(quantity), fmt.Sprint(product))["message"]
		response = append(response, new_response_item)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
