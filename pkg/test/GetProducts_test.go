package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	// "github.com/sanjana-jadhav-searce/shopping-catlog/pkg/controllers"
)

func GetProducts(t *testing.T) any {
	response, err := http.Get(URL + "/products")
	if err != nil {
		t.Log(err)
	}

	var v any

	err = json.NewDecoder(response.Body).Decode(&v)
	if err != nil {
		t.Log(err)
	}

	return v
}

// var URL string = "http://127.0.0.1:8000"

func TestGetProducts(t *testing.T) {
	products := GetProducts(t)

	_, ok := products.([]any)

	if !ok {
		t.Errorf("Expected a slice of products but got: " + fmt.Sprint(products))
	}

	products = GetCategory(t)

	_, ok = products.([]any)

	if !ok {
		t.Errorf("Expected an error of Products but got: " + fmt.Sprint(products))
	}
}
