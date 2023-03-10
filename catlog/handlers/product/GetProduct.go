package handlers_product

import (
	"net/http"

	"demo/helpers"
	"demo/queryhelpers"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	product_id := helpers.ParseMuxVarToInt(r, "id")
	product, err := queryhelpers.GetProduct(product_id)
	helpers.HandleError("GetProductQueryHelperError", err)

	if product != nil {
		helpers.SendResponse(product, w)
	} else {
		helpers.SendResponse(map[string]string{"message": "Product not found"}, w)
	}
}
