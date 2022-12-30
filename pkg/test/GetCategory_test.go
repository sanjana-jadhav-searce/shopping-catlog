package tests

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
	// "github.com/subramanyam-searce/product-catalog-go/helpers"
)

func GetcategoryViaAPI(name string, t *testing.T) map[string]string {
	response, err := http.Get(URL + "/productinventory" + fmt.Sprint(name))
	if err != nil {
		log.Panic(err)
	}

	response_json := map[string]string{}
	json.NewDecoder(response.Body).Decode(&response_json)

	return response_json
}

func TestGetCategory(t *testing.T) {

	// Valid product_id
	name := "Clothing"
	response := GetcategoryViaAPI(name, t)
	_, ok := response["Clothing"]
	dettol := map[string]any{

		"categoryname": "Clothing",
	}
	demo := map[string]any{
		"categoryname": "",
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
