package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"demo/helpers"
)

func TestCreateCart(t *testing.T) {
	response, err := http.Post("http://localhost:8000"+"/cart/create", "application/json", nil)
	helpers.HandleTestError(err, t)

	json_response := map[string]string{}
	json.NewDecoder(response.Body).Decode(&json_response)

	got_response, ok := json_response["ref"]

	if !ok {
		t.Errorf("Expected: %v, Got: %v", "{\"red\":\"<Unique ID here>\"}", got_response)
	}
}
