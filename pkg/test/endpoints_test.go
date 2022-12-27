package main

import (
	// "bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func GetProducts(t *testing.T) map[string]string {
	response, err := http.GET(URL + "/products")
	if err != nil {
		return err
	}

	response_json := map[string]string{}
	json.NewDecoder(response.Body).Decode(&response_json)
}

func TestGetProducts(t *testing.T) {
	response, err := GetProducts(t)
	if err != nil {
		t.Errorf("Expected Response : %v , Got Response : %v", "A valid product key", response)
	}
}
