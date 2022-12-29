package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	// "github.com/sanjana-jadhav-searce/shopping-catlog/pkg/controllers"
)

func GetInventory(t *testing.T) any {
	response, err := http.Get(URL + "/productsinventory")
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

func TestGetInventory(t *testing.T) {
	inventory := GetInventory(t)

	_, ok := inventory.([]any)

	if !ok {
		t.Errorf("Expected a slice of Inventory but got: " + fmt.Sprint(inventory))
	}

	inventory = GetInventory(t)

	_, ok = inventory.([]any)

	if !ok {
		t.Errorf("Expected an error of categories but got: " + fmt.Sprint(inventory))
	}
}
