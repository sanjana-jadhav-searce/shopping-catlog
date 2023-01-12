package helpers

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:Sanju2001@/shoppingcatlog?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Database Error:", err)
	}
}

func ConnectToDB() *sql.DB {
	return db
}
