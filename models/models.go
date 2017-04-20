package models

import (
  "github.com/astaxie/beego/orm"
)

type User struct {
  ID       int `orm:"column(id)"`
  UserName string `orm:"size(30)"`
  Email string `orm:"size(200);unique"`
  CreatedAt time.Time `orm:"auto_now_add;type"`
  UpdatedAt time.Time `orm:"auto_now"`
  Reviews []*Review `orm:"reverse(many)"`
}

type Review struct {
	ID      int    `orm:"column(id)"`
  User    *User `orm:"rel(fk)"`
	Title   string `orm:"size(250)"`
	Rating  int16
	Comment string  `orm:"type(text)"`
  CreatedAt time.Time `orm:"auto_now_add"`
  UpdatedAt time.Time `orm:"auto_now"`
}

func init() {
  orm.RegisterModel(new(User), new(Review))
}
