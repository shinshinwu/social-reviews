package controllers

import (
	"github.com/astaxie/beego"
  "social-reviews/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) ShowReviews() {
  user, err := models.FindUser(1)
  if err != nil {
    beego.BeeLogger.Error("Error finding first user%v", err.Error())
  }
	c.Data["PageTitle"] = "Review Index Page"
    c.Data["UserName"] = user.UserName
    c.Data["Email"] = user.Email
	c.TplName = "reviews.tpl"
}

func (c *MainController) CreateReview() {
  email := c.GetString("user[email]")
  username := c.GetString("user[username]")

  user, err := models.CreateUser(username, email)
  if err != nil {
    c.Ctx.WriteString("failed to create user")
  } else {
    c.Ctx.WriteString(user.Email)
    c.Ctx.WriteString(user.UserName)
  }
}
