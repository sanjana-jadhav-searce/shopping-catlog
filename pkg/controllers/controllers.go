package controllers

import (
	"database/sql"
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

var db *sql.DB

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
	x := "Inserted data successfully!"
	json.Marshal(x)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(x)

}

func GetProductByName(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()
	var product models.Product
	params := mux.Vars(r)
	result, err := db.Query("SELECT id, name, specification, SKU, category, price FROM products WHERE name=?", params["name"])
	fmt.Println(result)
	if err != nil {
		panic(err.Error())
	}
	for result.Next() {
		err := result.Scan(&product.ID, &product.Name, &product.Specification, &product.SKU, &product.Category, &product.Price)
		if err != nil {
			panic(err.Error())
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products models.MinimumProduct
	var product []models.MinimumProduct
	db := config.Connect()
	defer db.Close()
	rows, err := db.Query("SELECT id,name,price FROM products limit 20")
	fmt.Println(rows)
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		err = rows.Scan(&products.ID, &products.Name, &products.Price)

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

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)

	if err != nil {
		panic(err)
	}

	name := r.FormValue("name")
	specification := r.FormValue("specification")
	SKU := r.FormValue("SKU")
	category := r.FormValue("category")
	price := r.FormValue("price")
	id := r.FormValue("id")

	_, err = db.Exec("UPDATE products SET name=?, specification=?, SKU=?, category=?, price=? WHERE id=?", name, specification, SKU, category, price, id)

	if err != nil {
		log.Print(err)
	}

	x := "Updated data successfully!"
	json.Marshal(x)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(x)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	product := r.FormValue("product")
	rows, err := db.Query("SELECT * FROM products WHERE name=?", product)
	if product == "" {

		y := "Data Not Found"
		json.Marshal(y)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(y)
		return
	}
	if rows.Next() == false {
		z := "Invalid Inventory Product Reference"
		json.Marshal(z)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(z)
		return

	}
	result, err := db.Exec("DELETE FROM products WHERE name=?", product)
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
			msg := "Product item deleted successfully"
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

// func CreateInventoryProduct(w http.ResponseWriter, r *http.Request) {
// 	db := config.Connect()
// 	defer db.Close()

// 	err := r.ParseMultipartForm(4096)
// 	if err != nil {
// 		panic(err)
// 	}

// 	productname := r.FormValue("productname")
// 	quantity := r.FormValue("quantity")

// 	_, err = db.Exec("INSERT INTO inventories(productname, quantity) VALUES(?, ?)", productname, quantity)

// 	if err != nil {
// 		log.Print(err)
// 		return
// 	}
// 	x := "Inserted product into inventories successfully!"
// 	json.Marshal(x)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	json.NewEncoder(w).Encode(x)

// }

// func GetProductByNameInInventory(w http.ResponseWriter, r *http.Request) {
// 	db := config.Connect()
// 	defer db.Close()
// 	var product models.Inventory
// 	params := mux.Vars(r)
// 	result, err := db.Query("SELECT productname, quantity FROM inventories WHERE productname=?", params["name"])
// 	fmt.Println(result)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	for result.Next() {
// 		err := result.Scan(&product.Product, &product.Quantity)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(product)
// }

// func GetProductsInventory(w http.ResponseWriter, r *http.Request) {
// 	var products models.Inventory
// 	var product []models.Inventory
// 	db := config.Connect()
// 	defer db.Close()
// 	rows, err := db.Query("SELECT productname, quantity FROM inventories")
// 	fmt.Println(rows)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	for rows.Next() {
// 		err = rows.Scan(&products.Product, &products.Quantity)

// 		if err != nil {
// 			fmt.Println(err.Error())
// 		} else {
// 			product = append(product, products)
// 		}
// 		fmt.Println(product)
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	json.NewEncoder(w).Encode(product)
// }

// func UpdateProductInventory(w http.ResponseWriter, r *http.Request) {
// 	db := config.Connect()
// 	defer db.Close()

// 	err := r.ParseMultipartForm(4096)

// 	if err != nil {
// 		panic(err)
// 	}

// 	productname := r.FormValue("productname")
// 	quantity := r.FormValue("quantity")

// 	_, err = db.Exec("UPDATE inventories SET quantity=? WHERE productname=?", quantity, productname)

// 	if err != nil {
// 		log.Print(err)
// 	}

// 	x := "Updated Inventory data successfully!"
// 	json.Marshal(x)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(x)
// }

// func DeleteProductInventory(w http.ResponseWriter, r *http.Request) {
// 	db := config.Connect()
// 	defer db.Close()

// 	product := r.FormValue("product")
// 	rows, err := db.Query("SELECT productname, quantity FROM inventories WHERE productname=?", product)
// 	if product == "" {

// 		y := "Data Not Found"
// 		json.Marshal(y)
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(y)
// 		return
// 	}
// 	if rows.Next() == false {
// 		z := "Invalid Inventory Product Reference"
// 		json.Marshal(z)
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(z)
// 		return

// 	}
// 	result, err := db.Exec("DELETE FROM inventories WHERE productname=?", product)
// 	if err != nil {
// 		log.Print(err)
// 		return
// 	}
// 	rows_affected, err := result.RowsAffected()
// 	if err != nil {
// 		log.Print(err)
// 		return
// 	}
// 	if rows_affected != 0 {
// 		if err != nil {
// 			log.Print(err)
// 			return
// 		} else {
// 			msg := "Inventory item deleted successfully"
// 			json.Marshal(msg)
// 			w.Header().Set("Content-Type", "application/json")
// 			json.NewEncoder(w).Encode(msg)
// 			return

// 		}
// 	} else {
// 		msg := "Product is not found!"
// 		json.Marshal(msg)
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(msg)
// 		return

// 	}
// }

// func CreateCategory(w http.ResponseWriter, r *http.Request) {
// 	db := config.Connect()
// 	defer db.Close()

// 	err := r.ParseMultipartForm(4096)
// 	if err != nil {
// 		panic(err)
// 	}

// 	name := r.FormValue("name")

// 	_, err = db.Exec("INSERT INTO categorymaster(name) VALUES(?)", name)

// 	if err != nil {
// 		log.Print(err)
// 		return
// 	}
// 	x := "Inserted Category Data Successfully!"
// 	json.Marshal(x)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	json.NewEncoder(w).Encode(x)

// }

// func GetCategoryByName(w http.ResponseWriter, r *http.Request) {
// 	db := config.Connect()
// 	defer db.Close()
// 	var category models.CategoryMaster
// 	params := mux.Vars(r)
// 	result, err := db.Query("SELECT name FROM categorymaster WHERE name=?", params["name"])
// 	fmt.Println(result)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	for result.Next() {
// 		err := result.Scan(&category.Name)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(category)
// }

// func GetCategories(w http.ResponseWriter, r *http.Request) {
// 	var categories models.CategoryMaster
// 	var category []models.CategoryMaster
// 	db := config.Connect()
// 	defer db.Close()
// 	rows, err := db.Query("SELECT name FROM categorymaster")
// 	fmt.Println(rows)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	for rows.Next() {
// 		err = rows.Scan(&categories.Name)

// 		if err != nil {
// 			fmt.Println(err.Error())
// 		} else {
// 			category = append(category, categories)
// 		}
// 		fmt.Println(category)
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	json.NewEncoder(w).Encode(category)
// }

// func DeleteCategory(w http.ResponseWriter, r *http.Request) {
// 	db := config.Connect()
// 	defer db.Close()

// 	product := r.FormValue("product")
// 	rows, err := db.Query("SELECT name FROM categorymaster")
// 	if product == "" {

// 		y := "Data Not Found"
// 		json.Marshal(y)
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(y)
// 		return
// 	}
// 	if rows.Next() == false {
// 		z := "Invalid Category Reference"
// 		json.Marshal(z)
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(z)
// 		return

// 	}
// 	result, err := db.Exec("DELETE FROM categorymaster WHERE product=?", product)
// 	if err != nil {
// 		log.Print(err)
// 		return
// 	}
// 	rows_affected, err := result.RowsAffected()
// 	if err != nil {
// 		log.Print(err)
// 		return
// 	}
// 	if rows_affected != 0 {
// 		if err != nil {
// 			log.Print(err)
// 			return
// 		} else {
// 			msg := "Category item deleted successfully"
// 			json.Marshal(msg)
// 			w.Header().Set("Content-Type", "application/json")
// 			json.NewEncoder(w).Encode(msg)
// 			return

// 		}
// 	} else {
// 		msg := "Product is not found!"
// 		json.Marshal(msg)
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(msg)
// 		return

// 	}
// }

// func AddItemToCart(w http.ResponseWriter, r *http.Request) {

// 	db := config.Connect()
// 	defer db.Close()
// 	var products models.Cart
// 	var cart []models.Cart
// 	var cartitems = []models.Item{}
// 	var addition models.Total
// 	var totalcartvalue int64
// 	fmt.Println(cartitems)
// 	product := r.FormValue("product")
// 	quantity := r.FormValue("quantity")
// 	rows, err := db.Query("SELECT productname, quantity FROM inventories WHERE productname=?", product)
// 	fmt.Println(rows)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	for rows.Next() {
// 		err = rows.Scan(&products.Product, &products.Quantity)
// 		cartquantity, err1 := strconv.ParseInt(quantity, 0, 0)
// 		if err1 != nil {
// 			fmt.Println("error while parsing")
// 		}
// 		if err != nil {
// 			fmt.Println(err.Error())
// 		} else {
// 			if products.Quantity >= cartquantity {
// 				_, err2 := db.Exec("INSERT INTO carts(product, quantity) VALUES(?, ?)", product, quantity)

// 				if err2 != nil {
// 					log.Print(err2)
// 					return
// 				}
// 				insertion := "Data Inserted Successfully!"
// 				json.Marshal(insertion)
// 				w.Header().Set("Content-Type", "application/json")
// 				json.NewEncoder(w).Encode(insertion)
// 				message := "All Cart Items ---> "
// 				json.Marshal(message)
// 				rows, err3 := db.Query("SELECT product, quantity FROM carts")
// 				if err3 != nil {
// 					fmt.Println(err3)
// 				}
// 				for rows.Next() {
// 					err3 = rows.Scan(&products.Product, &products.Quantity)
// 					if err3 != nil {
// 						fmt.Println(err3.Error())
// 					} else {
// 						cart = append(cart, products)
// 					}
// 					fmt.Println(cart)
// 				}
// 				rows, err4 := db.Query("SELECT products.name,carts.quantity,products.price FROM products Cross JOIN carts ON products.name = carts.product ORDER BY products.name;")
// 				fmt.Println(rows)
// 				if err4 != nil {
// 					fmt.Println(err4)
// 				}

// 				for rows.Next() {
// 					err4 = rows.Scan(&addition.Product, &addition.Quantity, &addition.Price)
// 					totalcartvalue += (addition.Quantity * addition.Price)
// 				}
// 				json.Marshal(totalcartvalue)
// 				message2 := "Total Cart Value ---> "
// 				json.NewEncoder(w).Encode(message)
// 				json.NewEncoder(w).Encode(cart)
// 				json.NewEncoder(w).Encode(message2)
// 				json.NewEncoder(w).Encode(totalcartvalue)
// 			} else {
// 				x := "Data not inserted"
// 				json.Marshal(x)
// 				w.Header().Set("Content-Type", "application/json")
// 				json.NewEncoder(w).Encode(x)
// 			}
// 		}
// 	}

// }

// func CreateCart(w http.ResponseWriter, r *http.Request) {

// 	db := config.Connect()
// 	defer db.Close()
// 	var products models.Cart
// 	// var cart []models.Cart
// 	var cartitems = []models.Item{}
// 	// var addition models.Total
// 	// var totalcartvalue int64
// 	fmt.Println(cartitems)
// 	product := r.FormValue("product")
// 	quantity := r.FormValue("quantity")
// 	rows, err := db.Query("SELECT productname, quantity FROM inventories WHERE productname=?", product)
// 	fmt.Println(rows)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	for rows.Next() {
// 		err = rows.Scan(&products.Product, &products.Quantity)
// 		cartquantity, err1 := strconv.ParseInt(quantity, 0, 0)
// 		if err1 != nil {
// 			fmt.Println("error while parsing")
// 		}
// 		if err != nil {
// 			fmt.Println(err.Error())
// 		} else {
// 			if products.Quantity >= cartquantity {
// 				_, err2 := db.Exec("INSERT INTO carts(product, quantity) VALUES(?, ?)", product, quantity)

// 				if err2 != nil {
// 					log.Print(err2)
// 					return
// 				}
// 				insertion := "Data Inserted Successfully!"
// 				json.Marshal(insertion)
// 				w.Header().Set("Content-Type", "application/json")
// 				json.NewEncoder(w).Encode(insertion)

// 			} else {
// 				x := "Data not inserted"
// 				json.Marshal(x)
// 				w.Header().Set("Content-Type", "application/json")
// 				json.NewEncoder(w).Encode(x)
// 			}
// 		}
// 	}

// }

// func GetCart(w http.ResponseWriter, r *http.Request) {

// 	db := config.Connect()
// 	defer db.Close()
// 	var products models.Cart
// 	var cart []models.Cart

// 	w.Header().Set("Content-Type", "application/json")
// 	message := "All Cart Items ---> "
// 	json.Marshal(message)
// 	rows, err3 := db.Query("SELECT product, quantity FROM carts")
// 	if err3 != nil {
// 		fmt.Println(err3)
// 	}
// 	for rows.Next() {
// 		err3 = rows.Scan(&products.Product, &products.Quantity)
// 		if err3 != nil {
// 			fmt.Println(err3.Error())
// 		} else {
// 			cart = append(cart, products)
// 		}
// 		fmt.Println(cart)
// 	}
// 	json.NewEncoder(w).Encode(message)
// 	json.NewEncoder(w).Encode(cart)
// }

// func DeleteCart(w http.ResponseWriter, r *http.Request) {
// 	db := config.Connect()
// 	defer db.Close()

// 	product := r.FormValue("product")
// 	rows, err := db.Query("SELECT product, quantity FROM carts WHERE product=?", product)
// 	if product == "" {

// 		y := "Data Not Found"
// 		json.Marshal(y)
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(y)
// 		return
// 	}
// 	if rows.Next() == false {
// 		z := "Invalid Cart Reference"
// 		json.Marshal(z)
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(z)
// 		return

// 	}
// 	result, err := db.Exec("DELETE FROM carts WHERE product=?", product)
// 	if err != nil {
// 		log.Print(err)
// 		return
// 	}
// 	rows_affected, err := result.RowsAffected()
// 	if err != nil {
// 		log.Print(err)
// 		return
// 	}
// 	if rows_affected != 0 {
// 		if err != nil {
// 			log.Print(err)
// 			return
// 		} else {
// 			msg := "Cart item deleted successfully"
// 			json.Marshal(msg)
// 			w.Header().Set("Content-Type", "application/json")
// 			json.NewEncoder(w).Encode(msg)
// 			return

// 		}
// 	} else {
// 		msg := "Product is not found!"
// 		json.Marshal(msg)
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(msg)
// 		return

// 	}
// }
