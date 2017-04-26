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

func (c *MainController) Reviews() {
  user, err := models.FindUser(1)
  if err != nil {
    beego.BeeLogger.Error("Error finding first user%v", err.Error())
  }
	c.Data["PageTitle"] = "Review Index Page"
    c.Data["UserName"] = user.UserName
    c.Data["Email"] = user.Emails
	c.TplName = "reviews.tpl"
}
