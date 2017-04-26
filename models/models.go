package models

import (
  "time"
  "social-reviews/db"
  "github.com/astaxie/beego"
)

type User struct {
  ID              int    `gorm:"primary_key"`
  UserName        string  `gorm:"size:255;not null;unique_index"` // Default size for string is 255, reset it with this tag
  Email           string  `gorm:"size:255;unique_index"`
  CreatedAt       time.Time
  UpdatedAt       time.Time

  Reviews         []Review
}

type Review struct {
	ID              int    `gorm:"primary_key"`
  UserID          int     `gorm:"index"`
	Title           string  `gorm:"not null"`
	Comment         string  `gorm:"type:text;not null"`
  IsApproved      bool    `gorm:"default:1"`
  CreatedAt       time.Time
  UpdatedAt       time.Time

  User            User
}

func init() {
  db.RegisterModel(&User{})
  db.RegisterModel(&Review{})
}

func AllUsers() (users []*User, err error) {
  res := db.Conn.Debug().Order("first_name asc, last_name asc").Find(&users)
  err = res.Error
  if err != nil {
    beego.BeeLogger.Error("Error finding users: %v", res.Error.Error())
  }
  return
}

func FindUser(id int) (*User, error) {
  user := &User{ID: id}
  res := db.Conn.Debug().First(&user, "id=?", id)
  err := res.Error
  if err != nil {
    beego.BeeLogger.Error("Error finding user with id %s: %v", id, res.Error.Error())
  }
  return user, err
}
