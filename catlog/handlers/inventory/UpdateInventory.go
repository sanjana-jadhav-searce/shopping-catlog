package handlers_inventory

import (
	"encoding/json"
	"net/http"

	"demo/helpers"
	"demo/typedefs"
)

func UpdateInventory(w http.ResponseWriter, r *http.Request) {
	inventory_item := typedefs.Inventory{}
	json.NewDecoder(r.Body).Decode(&inventory_item)

	rows, err := helpers.RunQuery("SELECT quantity FROM inventory WHERE product_id=?", inventory_item.ProductID)
	helpers.HandleError("runQueryError", err)

	if rows.Next() {
		var existing_quantity int
		err = rows.Scan(&existing_quantity)
		helpers.HandleError("rowsScanError", err)

		_, err = helpers.RunQuery("UPDATE inventory SET quantity=? WHERE product_id=?", existing_quantity+inventory_item.Quantity, inventory_item.ProductID)
		helpers.HandleError("runQueryError", err)

		if err != nil {
			helpers.SendResponse(map[string]string{"message": err.Error()}, w)
			return
		}

		if existing_quantity+inventory_item.Quantity == 0 {
			_, err = helpers.RunQuery("DELETE FROM inventory WHERE product_id=?", inventory_item.ProductID)
			helpers.HandleError("runQueryError", err)

			if err != nil {
				helpers.SendResponse(map[string]string{"message": err.Error()}, w)
				return
			}
		}
	} else {
		query := "INSERT INTO inventory VALUES(?, ?)"
		_, err = helpers.RunQuery(query, inventory_item.ProductID, inventory_item.Quantity)
		helpers.HandleError("runQueryError", err)
		if err != nil {
			helpers.SendResponse(map[string]string{"message": err.Error()}, w)
			return
		}
	}

	helpers.SendResponse(map[string]string{"message": "Inventory Items added successfully"}, w)
}
