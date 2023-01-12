package queryhelpers

import (
	"demo/helpers"
	"demo/typedefs"
)

func GetInventory(id int) (*typedefs.Inventory, error) {
	query := "SELECT * FROM inventory WHERE product_id=?;"
	var inventoryItem *typedefs.Inventory = nil

	rows, err := helpers.RunQuery(query, id)

	if rows.Next() {
		inventoryItem = &typedefs.Inventory{}
		rows.Scan(&inventoryItem.ProductID, &inventoryItem.Quantity)
	}

	return inventoryItem, err
}
