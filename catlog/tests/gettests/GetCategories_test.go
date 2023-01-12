package tests

import (
	"io/ioutil"
	"net/http"
	"testing"

	"demo/helpers"
)

func TestGetCategory(t *testing.T) {

	resp, err := http.Get("http://localhost:8000/categories")
	if err != nil {
		t.Errorf("Error making request: %v", err)
		helpers.HandleTestError(err, t)
	}

	_, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Errorf("Error reading response body: %v", err)
		helpers.HandleTestError(err, t)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

}
