package tests

import (
	"os"
	"testing"

	"demo/helpers"
)

func RestoreDBTestingState(t *testing.T) {
	db := helpers.ConnectToDB()

	sqlString, err := os.ReadFile("../sql_commands/testing_state.sql")
	helpers.HandleTestError("readFileError", err, t)

	_, err = db.Exec(string(sqlString))
	helpers.HandleTestError("dbExecError", err, t)
}
