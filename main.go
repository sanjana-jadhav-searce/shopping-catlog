package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/sanjana-jadhav-searce/shopping-catlog/pkg/routes"
)

// var db *sql.DB
// var err error

func main() {

	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	r := mux.NewRouter()
	routes.RegisterShoppingCart(r)
	http.Handle("/", r)
	fmt.Println("Server started on PORT 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
