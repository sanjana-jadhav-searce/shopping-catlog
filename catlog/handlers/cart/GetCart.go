package handlers_cart

import (
	"net/http"

	"demo/helpers"
	"demo/typedefs"
)

func GetCart(w http.ResponseWriter, r *http.Request) {
	urlQuery := r.URL.Query()
	ref := urlQuery.Get("ref")

	rows, err := helpers.RunQuery("SELECT * FROM cart_reference WHERE ref=?", ref)
	helpers.HandleError("runQuery", err)

	if !rows.Next() {
		helpers.SendResponse(map[string]string{"message": "Cart Reference is Invalid"}, w)
		return
	}

	cart := typedefs.Cart{}
	err = rows.Scan(&cart.Ref, &cart.CreatedAt)
	helpers.HandleError("rowsScanError", err)

	rows, err = helpers.RunQuery("SELECT * FROM cart_item WHERE ref=?", ref)
	helpers.HandleError("runQueryError", err)

	for rows.Next() {
		new_cart_item := typedefs.CartItem{}
		cart_ref_id := ""
		err = rows.Scan(&cart_ref_id, &new_cart_item.ProductID, &new_cart_item.Quantity)
		helpers.HandleError("rowsScanError", err)
		cart.Items = append(cart.Items, new_cart_item)
	}

	cart.EvaluateCartValue()
	helpers.SendResponse(cart, w)

}
