package ConsoleInterface

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func Product() {
	fmt.Println("Welcome to our products section,You can perform CRUD operations on 'Product' table")
	fmt.Printf("1.Add\n2.Get product details with product id\n3.Get products with minimum details\n4.Update\n5.Delete\n")
	fmt.Println("Please enter your choice")
	var choice int
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		fmt.Println(err)
	}
	if choice == 1 {
		AddProduct()
	} else if choice == 2 {
		GetProduct()
	} else if choice == 3 {
		GetProducts()
	} else if choice == 4 {
		UpdateProduct()
	} else if choice == 5 {
		DeleteProduct()
	}
}

func AddProduct() {
	fmt.Println("Please enter the valid product id")
	var product_id int
	_, err := fmt.Scanf("%d", &product_id)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Please enter the product name")
	var name string
	_, err = fmt.Scanln(&name)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Please enter the specification key")
	var key string
	_, err = fmt.Scanln(&key)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Please enter the specification value")
	var value string
	_, err = fmt.Scanln(&value)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Please enter the SKU")
	var sku int
	_, err = fmt.Scanf("%d", &sku)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Please enter the Category id")
	var category_id int
	_, err = fmt.Scanf("%d", &category_id)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Please enter the Price")
	var price float64
	_, err = fmt.Scanf("%f", &price)
	if err != nil {
		fmt.Println(err)
	}

	own_data := fmt.Sprintf("{\"product_id\":%v,\"product_name\":\"%v\",\"specification\": {\"%v\":\"%v\"},\"sku\":\"%v\",\"category_id\":%v,\"price\":%v}", product_id, name, key, value, sku, category_id, price)

	data := []byte(own_data)

	_, err = http.Post("http://localhost:8000/add/product", "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Do you want to continue? (y/n)")
	var opt string
	_, err = fmt.Scanln(&opt)
	if err != nil {
		fmt.Println(err)
	}
	if opt == "y" {
		Console()
	} else {
		return
	}

}

func GetProduct() {
	fmt.Println("Please enter product id")
	var product_id int
	_, err := fmt.Scanln(&product_id)
	if err != nil {
		fmt.Println(err)
	}

	_, err = http.Get("http://localhost:8000/product/" + fmt.Sprint(product_id))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Do you want to continue? (y/n)")
	var opt string
	_, err = fmt.Scanln(&opt)
	if err != nil {
		fmt.Println(err)
	}
	if opt == "y" {
		Console()
	} else {
		return
	}

}

func GetProducts() {
	fmt.Println("Please enter page no. to get the product")
	var page string
	_, err := fmt.Scanln(&page)
	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:8000/products?page=%v", page), nil)
	if err != nil {
		fmt.Println(err)
	}

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Do you want to continue? (y/n)")
	var opt string
	_, err = fmt.Scanln(&opt)
	if err != nil {
		fmt.Println(err)
	}
	if opt == "y" {
		Console()
	} else {
		return
	}
}

func UpdateProduct() {
	fmt.Println("Please enter the product id")
	var product_id string
	_, err := fmt.Scanln(&product_id)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Please enter the key to be updated")
	var key string
	_, err = fmt.Scanln(&key)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Please enter the value to be updated")
	var value string
	_, err = fmt.Scanln(&value)
	if err != nil {
		fmt.Println(err)
	}

	data := map[string]any{key: value}
	my_data, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	request_body := bytes.NewBuffer(my_data)
	req, err := http.NewRequest("POST", fmt.Sprintf("http://localhost:8000/product/update/%v", product_id), request_body)
	if err != nil {
		fmt.Println(err)
	}

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Update done succesfully")

	fmt.Println("Do you want to continue? (y/n)")
	var opt string
	_, err = fmt.Scanln(&opt)
	if err != nil {
		fmt.Println(err)
	}
	if opt == "y" {
		Console()
	} else {
		return
	}
}

func DeleteProduct() {
	fmt.Println("Please enter the product id")
	var product_id string
	_, err := fmt.Scanln(&product_id)
	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://localhost:8000/product/delete/%v", product_id), nil)
	if err != nil {
		fmt.Println(err)
	}

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Delete done succesfully")

	fmt.Println("Do you want to continue? (y/n)")
	var opt string
	_, err = fmt.Scanln(&opt)
	if err != nil {
		fmt.Println(err)
	}
	if opt == "y" {
		Console()
	} else {
		return
	}

}
