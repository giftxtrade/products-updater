package main

import (
	"fmt"

	"github.com/giftxtrade/products-updater/src/config"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db_config, err := config.LoadConfig()
	if err != nil {
		panic(err.Error())
	}

	db := config.Connect(db_config["user"], db_config["pass"], db_config["db_name"])
	defer db.Close()


	rows, err := db.Query("SELECT productKey FROM products")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var keys []string = make([]string, 1000)
	for rows.Next() {
		var key string
		if err = rows.Scan(&key); err != nil {
			panic(err.Error())
		}
		keys = append(keys, key)
	}
	fmt.Println(keys)
}