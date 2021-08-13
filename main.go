package main

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	args := os.Args
	
	if len(args) < 3 {
		panic("Provide username and password")
	}
	username := args[1]
	password := args[2]
	db_name := "giftxtrade2"

	db, err := sql.Open("mysql", username + ":" + password + "@tcp(127.0.0.1:3306)/" + db_name)
    if err != nil {
        panic(err.Error())
    }

    // defer the close till after the main function has finished executing
    defer db.Close()
}