package config

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

type Config struct {
  SigningKey string
}
func ConnectToDB() *sql.DB {
  db, err := sql.Open("mysql", "homestead:secret@tcp(192.168.10.10:3306)/schools")
  // if there is an error opening the connection, handle it
  if err != nil {
    panic(err.Error())
  }

  return db
}

func GetConfig() Config {
  var config Config
  config.SigningKey = "captainjacksparrowsayshi"
  return config
}


