package tests

import (
	"log"
	"testing"

	"github.com/sanjana-jadhav-searce/shopping-catlog/pkg/config"
)

func GetConnection(t *testing.T) {

	db := config.GetDB()
	log.Print(db)

}
