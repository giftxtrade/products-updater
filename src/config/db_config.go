package config

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
)

func Connect(user string, pass string, db_name string) *sql.DB {
	url := user + ":" + pass + "@tcp(localhost:3306)/" + db_name

	var db *sql.DB
	var err error
	if db, err = sql.Open("mysql", url); err != nil {
		panic("Could not connect to db.\n" + err.Error())
	}
	return db
}

func LoadConfig() (map[string]string, error) {
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