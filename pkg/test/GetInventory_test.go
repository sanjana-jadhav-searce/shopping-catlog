package tests

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
	// "github.com/subramanyam-searce/product-catalog-go/helpers"
)

func GetInventoryViaAPI(name string, t *testing.T) map[string]string {
	response, err := http.Get(URL + "/productinventory" + fmt.Sprint(name))
	if err != nil {
		log.Panic(err)
	}

	response_json := map[string]string{}
	json.NewDecoder(response.Body).Decode(&response_json)

	return response_json
}

func TestGetInventories(t *testing.T) {

	// Valid product_id
	name := "Dettol"
	response := GetInventoryViaAPI(name, t)
	_, ok := response["Dettol"]
	dettol := map[string]any{

		"productname": "Dettol",
		"quantity":    1000,
	}
	demo := map[string]any{
		"productname": "",
		"quantity":    0,
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
