package helpers

import (
	"testing"
)

func HandleTestError(err error, t *testing.T) {
	if err != nil {
		t.Log(err)
	}
}
