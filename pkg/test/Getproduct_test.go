package tests

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
	// "github.com/subramanyam-searce/product-catalog-go/helpers"
)

func GetProductViaAPI(name string, t *testing.T) map[string]string {
	response, err := http.Get(URL + "/product/" + fmt.Sprint(name))
	if err != nil {
		log.Panic(err)
	}

	response_json := map[string]string{}
	json.NewDecoder(response.Body).Decode(&response_json)

	return response_json
}

func TestGetProduct(t *testing.T) {

	// Valid product_id
	name := "Dettol"
	response := GetProductViaAPI(name, t)
	_, ok := response["Dettol"]
	demo := map[string]any{
		"id":            "",
		"name":          "",
		"specification": "",
		"SKU":           "",
		"category":      "",
		"price":         "",
	}
	dettol := map[string]any{
		"id":            2,
		"name":          "Dettol",
		"specification": "Handwash",
		"SKU":           "50",
		"category":      "Essentials",
		"price":         65001,
	}
	if ok {
		t.Errorf("Expected Response: %v, Got Response: %v", dettol, response)
	}

	name = "helicopter"
	response = GetProductViaAPI(name, t)
	_, ok = response["message"]
	if ok {
		t.Errorf("Expected Response: %v, Got Response: %v", demo, response)
	}
}
