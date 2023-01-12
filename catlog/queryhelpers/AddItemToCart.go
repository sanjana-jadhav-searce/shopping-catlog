package queryhelpers

import (
	"fmt"
	"strconv"

	"demo/helpers"
	"demo/typedefs"
)

func AddItemToCart(ref, quantity_str, product_id string) map[string]string {

	if ref == "" || quantity_str == "" || product_id == "" {
		return map[string]string{"message": "ref / quantity / product_id missing in the url"}
	}

	quantity, err := strconv.Atoi(quantity_str)
	helpers.HandleError("strconvError:", err)

	if quantity < 0 {
		return map[string]string{"message": "Quantity cannot be Negative"}
	}

	rows, err := helpers.RunQuery("SELECT * FROM cart_reference WHERE ref=?;", ref)
	helpers.HandleError("runQueryError:", err)

	if !rows.Next() {
		return map[string]string{"message": "Invalid cart_reference"}
	}

	rows, err = helpers.RunQuery("SELECT p.product_id, i.quantity, p.name FROM product p LEFT JOIN inventory i ON p.product_id=i.product_id WHERE p.product_id=?", product_id)
	helpers.HandleError("runQueryError", err)

	if !rows.Next() {
		return map[string]string{"message": "Product id is invalid"}
	}

	inventory_item := typedefs.ProductInventory{}
	err = rows.Scan(&inventory_item.ProductID, &inventory_item.Quantity, &inventory_item.ProductName)
	helpers.HandleError("rowsScanError", err)

	if inventory_item.Quantity-quantity < 0 {
		return map[string]string{"message": "The Required Quantity is more than the Available Inventory Quantity: " + fmt.Sprint(inventory_item.Quantity)}
	}

	if inventory_item.Quantity-quantity > 0 {
		_, err = helpers.RunQuery("UPDATE inventory SET quantity=? WHERE product_id=?", inventory_item.Quantity-quantity, product_id)
		helpers.HandleError("runQueryError:", err)
	}

	if inventory_item.Quantity-quantity == 0 {
		_, err = helpers.RunQuery("DELETE FROM inventory WHERE product_id=?", product_id)
		helpers.HandleError("runQueryError:", err)
	}

	rows, err = helpers.RunQuery("SELECT quantity FROM cart_item WHERE ref=? AND product_id=?", ref, product_id)
	helpers.HandleError("runQueryError:", err)

	if rows.Next() {
		var db_quantity int
		rows.Scan(&db_quantity)

		_, err = helpers.RunQuery("UPDATE cart_item SET quantity=? WHERE ref=? AND product_id=?", db_quantity+quantity, ref, product_id)
		helpers.HandleError("runQueryError:", err)

	} else {
		_, err = helpers.RunQuery("INSERT INTO cart_item VALUES(?, ?, ?);", ref, product_id, quantity)
		helpers.HandleError("runQueryError:", err)
	}

	if err != nil {
		return map[string]string{"message": err.Error()}
	}

	return map[string]string{"message": "Item was added to the cart"}
}
