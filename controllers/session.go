package controllers

import (
  "github.com/astaxie/beego"
  "html/template"
  "social-reviews/models"
)

type SessionController struct {
  beego.Controller
}

func (c *SessionController) Prepare() {
  c.Layout = "layout/application.tpl"
  c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
}

func (c *SessionController) New() {
  c.TplName = "session/new.tpl"
  c.Data["PageTitle"] = "Login Page"
}

func (c *SessionController) Create() {
  email := c.GetString("email")
  password := c.GetString("password")
  user, err := models.UserLogin(email, password)
  if err != nil {
    c.Redirect("/login", 302)
  } else {
    c.SetSession("current_user", user.AsMap())
    c.Redirect("/reviews", 302)
  }
}

func (c *SessionController) Destroy() {
  c.DelSession("current_user")
  c.DestroySession()
  c.Redirect("/login", 302)
}
