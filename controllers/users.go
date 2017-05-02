package controllers

import (
  "github.com/astaxie/beego"
  "html/template"
  "social-reviews/models"
)

type UsersController struct {
  beego.Controller
}

func (c *UsersController) Prepare() {
  c.Layout = "layout/application.tpl"
  c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
}

func (c *UsersController) New() {
  c.Data["PageTitle"] = "Login Page"
  c.TplName = "users/new.tpl"
}

func (c *UsersController) Create() {
  email := c.GetString("email")
  password := c.GetString("password")
  username := c.GetString("username")
  user, err := models.CreateUser(username, email, password)
  if err != nil {
    c.Redirect("/login", 302)
  } else {
    c.SetSession("current_user", user.AsMap())
    c.Redirect("/reviews", 302)
  }
}
