package handlers_cart

import (
	"net/http"
	"time"

	"demo/helpers"

	"github.com/google/uuid"
)

func CreateCart(w http.ResponseWriter, r *http.Request) {
	ref := uuid.New()

	_, err := helpers.RunQuery("INSERT INTO cart_reference VALUES(?, ?);", ref, time.Now())
	helpers.HandleError("runQueryError:", err)

	if err != nil {
		helpers.SendResponse(map[string]string{"message": err.Error()}, w)
	} else {
		helpers.SendResponse(map[string]uuid.UUID{"ref": ref}, w)
	}

}
