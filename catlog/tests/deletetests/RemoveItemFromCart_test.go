package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"demo/helpers"
)

func RemoveItemFromCartViaAPI(ref string, product_id int, t *testing.T) map[string]string {
	request_url := fmt.Sprintf("%v/removeitemfromcart?ref=%v&product_id=%v", URL, ref, product_id)

	req, err := http.NewRequest("DELETE", request_url, nil)
	helpers.HandleTestError(err, t)

	res, err := http.DefaultClient.Do(req)
	helpers.HandleTestError(err, t)

	var v map[string]string
	json.NewDecoder(res.Body).Decode(&v)

	return v
}

func TestRemoveItemFromCart(t *testing.T) {
	test_cases := []map[string]any{
		//Valid ref and valid product id
		{"ref": "4d4d8297-7663-451d-b79e-49a545728552", "product_id": 3, "expected_response": "All quantity of this product was removed from your cart"},

		//invalid ref
		{"ref": "abcd", "product_id": 3, "expected_response": "Invalid cart_reference"},

		//valid ref and valid product id not present in cart
		{"ref": "4d4d8297-7663-451d-b79e-49a545728552", "product_id": 2, "expected_response": "Product is not found in your cart"},

		//valid ref and invalid product id
		{"ref": "4d4d8297-7663-451d-b79e-49a545728552", "product_id": 540, "expected_response": "Product is not found in your cart"},
	}

	for _, v := range test_cases {
		response := RemoveItemFromCartViaAPI(v["ref"].(string), v["product_id"].(int), t)

		if response["message"] != v["expected_response"] {
			t.Errorf("Expected: %v, Got: %v", v["expected_response"], response["message"])
		}
	}
}
