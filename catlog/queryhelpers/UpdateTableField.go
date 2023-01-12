package queryhelpers

import (
	"demo/helpers"
)

func UpdateTableField(tableName string, condition string, fieldName string, val string) error {
	query := "UPDATE " + tableName + " SET " + fieldName + "='" + val + "' " + condition + ";"
	_, err := helpers.RunQuery(query)

	return err
}
