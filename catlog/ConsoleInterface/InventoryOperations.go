package ConsoleInterface

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func Inventory() {
	fmt.Println("Welcome to our inventory section,please perform CRUD operations on 'Inventory' table")
	fmt.Printf("1.Add\n2.Get\n3.Update\n4.Delete\n")
	fmt.Println("Please enter your choice")
	var choice int
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		fmt.Println(err)
	}
	if choice == 1 {
		GetInventory()
	} else if choice == 2 {
		UpdateInventory()
	}
}

func GetInventory() {

	_, err := http.Get("http://localhost:8000/inventory")
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

func UpdateInventory() {
	fmt.Println("Please enter the product id")
	var product_id string
	_, err := fmt.Scanln(&product_id)
	if err != nil {
		fmt.Println(err)
	}

	key := "quantity"

	fmt.Println("Please enter the quantity to be updated")
	var value int
	_, err = fmt.Scanln(&value)
	if err != nil {
		fmt.Println(err)
	}

	data := map[string]any{key: value}
	byte_data, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	request_body := bytes.NewBuffer(byte_data)
	req, err := http.NewRequest("PUT", "http://localhost:8000/inventory/update", request_body)
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
