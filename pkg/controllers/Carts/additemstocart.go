package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func AddItemsToCart(w http.ResponseWriter, r *http.Request) {
	response := []map[string]any{}
	request_body := []map[string]any{}

	err := json.NewDecoder(r.Body).Decode(&request_body)
	if err != nil {
		fmt.Println("ERROR IN DECODING", err)
	}
	fmt.Println(request_body)
	for _, v := range request_body {
		new_response_item := map[string]any{}
		product := v["product"]
		quantity := v["quantity"]

		url := "http://127.0.0.1:8000/additemstocart" + "&product=" + fmt.Sprint(product) + "&quantity=" + fmt.Sprint(quantity)
		fmt.Println(url)
		_, err = http.Post(url, "application/json", nil)
		if err != nil {
			fmt.Println("ERROR IN DECODING", err)
		}

		new_response_item["product"] = product
		new_response_item["quantity"] = quantity
		new_response_item["message"] = MultipleCart(fmt.Sprint(quantity), fmt.Sprint(product), w, r)["message"]
		response = append(response, new_response_item)

	}

	fmt.Println("response", response)
	json.NewEncoder(w).Encode(response)

}
