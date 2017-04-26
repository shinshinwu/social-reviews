package models

import (
  "social-reviews/db"
  "github.com/astaxie/beego"
)

type User struct {
  ID               int
  UserName         string  `gorm:"size:255;not null;unique"` // Default size for string is 255, reset it with this tag
  Emails           string  `gorm:"size:255;unique"`
}

// type Review struct {
// 	ID      int    `orm:"column(id)"`
//   User    *User `orm:"rel(fk)"`
// 	Title   string `orm:"size(250)"`
// 	Rating  int16
// 	Comment string  `orm:"type(text)"`
//   CreatedAt time.Time `orm:"auto_now_add"`
//   UpdatedAt time.Time `orm:"auto_now"`
// }

func init() {
  db.RegisterModel(&User{})
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
