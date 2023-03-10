package handlers_category

import (
	"net/http"

	"demo/helpers"
	"demo/typedefs"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	categories := []typedefs.Category{}
	var err error

	query := "SELECT * FROM category;"
	rows, err := helpers.RunQuery(query)
	helpers.HandleError("runQueryError", err)

	for rows.Next() {
		category := typedefs.Category{}
		err := rows.Scan(&category.CategoryID, &category.Name)
		helpers.HandleError("rowsScanError", err)
		categories = append(categories, category)
	}

	helpers.SendResponse(categories, w)
}
