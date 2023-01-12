package handlers_category

import (
	"encoding/json"
	"net/http"

	"demo/helpers"
	"demo/typedefs"
)

func AddCategory(w http.ResponseWriter, r *http.Request) {
	category := typedefs.Category{}
	json.NewDecoder(r.Body).Decode(&category)

	query := "INSERT INTO category VALUES(?, ?)"
	_, err := helpers.RunQuery(query, category.CategoryID, category.Name)
	helpers.HandleError("runQueryError", err)

	if err != nil {
		helpers.SendResponse(map[string]string{"message": err.Error()}, w)
	} else {
		helpers.SendResponse(map[string]string{"message": "Category added successfully"}, w)
	}
}
