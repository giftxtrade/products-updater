package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
)

func connect_db(user string, pass string, db_name string) *sql.DB {
	url := user + ":" + pass + "@tcp(localhost:3306)/" + db_name

	var db *sql.DB
	var err error
    if db, err = sql.Open("mysql", url); err != nil {
        panic("Could not connect to db.\n" + err.Error())
    }
	return db
}

func load_db_config() (map[string]string, error) {
	db_config_file_data, err := ioutil.ReadFile("db_config.json")
	if err != nil {
		return map[string]string{}, errors.New("db_config.json file not found")
	}

	var db_config map[string]string
	err = json.Unmarshal([]byte(db_config_file_data), &db_config)
	if err != nil {
		return map[string]string{}, err
	}
	return db_config, nil
}

func main() {
	db_config, err := load_db_config()
	if err != nil {
		panic(err.Error())
	}

	db := connect_db(db_config["user"], db_config["pass"], db_config["db_name"])
    defer db.Close()


	rows, err := db.Query("SELECT productKey from products")
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