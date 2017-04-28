package db

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

var Conn *gorm.DB

func Connect() (db *gorm.DB, err error) {
  db, err = gorm.Open("mysql", "root:@/social_review?parseTime=true")
  Conn = db
  return
}
