package main

import (
  _ "social-reviews/routers"
  _ "social-reviews/models" // import models package so models initialize and register
  "github.com/astaxie/beego"
  "social-reviews/db"
  "github.com/jinzhu/gorm"
)

func main() {
  conn := getDBConnection()
  defer conn.Close()
  beego.Run()
}

func getDBConnection() *gorm.DB {
  conn, err := db.Connect()
  if err != nil {
    panic(err.Error())
  }
  db.AutoMigrate()
  return conn
}

