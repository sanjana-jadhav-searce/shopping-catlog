package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"demo/helpers"
	"demo/typedefs"
)

func AddProductViaAPI(product typedefs.Product) string {
	json_product, err := json.Marshal(product)
	helpers.HandleError("jsonMarshalError", err)

	request_body := bytes.NewBuffer(json_product)

	res, err := http.Post("http://localhost:8000"+"/product/add", "application/json", request_body)
	helpers.HandleError("httpPostRequestError", err)

	var json_response map[string]string
	json.NewDecoder(res.Body).Decode(&json_response)

	response_message := json_response["message"]
	return response_message
}

func checkAddProductResponse(t *testing.T, expected_response string, got_response string) {
	if got_response != expected_response {
		t.Errorf("Expected Response: %v, Got Response: %v", expected_response, got_response)
	}

}

func TestAddProduct(t *testing.T) {
	product := typedefs.Product{
		Product_ID: 455,
		Name:       "Shorts",
		Specification: map[string]string{
			"color": "Brown",
			"size":  "40",
		},
		SKU:        "65548",
		CategoryID: 1,
		Price:      9.99,
	}

	got_response := AddProductViaAPI(product)
	expected_response := "Product Added Successfully"
	checkAddProductResponse(t, expected_response, got_response)

	product.Product_ID = 1
	got_response = AddProductViaAPI(product)
	expected_response = "pq: duplicate key value violates unique constraint \"product_pkey\""
	checkAddProductResponse(t, expected_response, got_response)

	product.Product_ID = 400
	product.CategoryID = 543
	got_response = AddProductViaAPI(product)
	expected_response = "pq: insert or update on table \"product\" violates foreign key constraint \"product_category_id_fkey\""
	checkAddProductResponse(t, expected_response, got_response)
}
