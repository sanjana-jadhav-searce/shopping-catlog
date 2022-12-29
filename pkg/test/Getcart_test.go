package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	// "github.com/sanjana-jadhav-searce/shopping-catlog/pkg/controllers"
)

func GetCart(t *testing.T) any {
	response, err := http.Get(URL + "/getcart")
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

func TestGetCart(t *testing.T) {
	carts := GetCategory(t)

	_, ok := carts.([]any)

	if !ok {
		t.Errorf("Expected a slice of categories but got: " + fmt.Sprint(carts))
	}

	carts = GetCategory(t)

	_, ok = carts.([]any)

	if !ok {
		t.Errorf("Expected an error of categories but got: " + fmt.Sprint(carts))
	}
}
