package main

import (
  "fmt"
	_ "social-reviews/routers"
	"github.com/astaxie/beego"
  _ "github.com/go-sql-driver/mysql"
)

func init() {
  orm.RegisterDriver("mysql", orm.DRMySQL)
  orm.RegisterDataBase("social-reviews", "mysql", "root:root@/social_reviews?charset=utf8")
}

func main() {
  o := orm.NewOrm()
  o.Using("social_reviews")

  user := new(User)
  user.UserName = "banana"
  user.Email = "banana@gmail.com"

  review := new(Review)
  review.User = user
  review.Title = "First Review"
  review.Rating = 4
  review.Comment = "Test test reviews!"

  fmt.Println(o.Insert(user))
  fmt.Println(o.Insert(review))

	beego.Run()
}
