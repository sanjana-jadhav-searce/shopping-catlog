package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestUpdateProduct(t *testing.T) {
	product_master := map[string]any{
		"id":            1,
		"name":          "Casio",
		"specification": "71",
		"sku":           "34a",
		"category_id":   1,
		"price":         2000,
	}

	CheckUpdateProduct(product_master, "", t)

	//updating only one field
	product_master = map[string]any{"id": 1, "name": "Casio", "price": 1090}
	CheckUpdateProduct(product_master, "", t)

	//product id not exists
	product_master = map[string]any{"id": 32, "name": "Jacket", "price": 1099}
	CheckUpdateProduct(product_master, "", t)

}

func CheckUpdateProduct(product_master map[string]any, response string, t *testing.T) {
	json_product, err := json.Marshal(product_master)
	if err != nil {
		fmt.Println("error", err)
	}
	request_body := bytes.NewBuffer(json_product)
	req, err := http.NewRequest("PUT", fmt.Sprintf("%v/updateproduct", URL), request_body)
	if err != nil {
		fmt.Println("error", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("error", err)
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error", err)
	}
	if string(bodyBytes) != response {
		t.Errorf("Expected: %v, Got: %v", response, string(bodyBytes))
	}
}
