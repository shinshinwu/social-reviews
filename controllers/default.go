package controllers

import (
  "encoding/json"
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
  reviews, err := models.AllReviews()
  if err != nil {
    beego.BeeLogger.Error("Error finding reviews %v", err.Error())
  }
	c.Data["PageTitle"] = "Review Index Page"
  c.Data["reviews"] = reviews
	c.TplName = "reviews.tpl"
}

func (c *MainController) CreateReview() {
  email := c.GetString("user[email]")
  username := c.GetString("user[username]")

  user, err := models.CreateUser(username, email)
  if err != nil {
    c.Ctx.WriteString("failed to create user")
  }

  title := c.GetString("review[title]")
  comment := c.GetString("review[comment]")

  review, err := models.CreateReview(user.ID, title, comment)
  if err != nil {
    c.Ctx.WriteString("failed to create review")
  }

  data, err := json.Marshal(review)
  if err != nil {
      panic(err)
  }

  c.Data["review"] = data
  c.ServeJSON()
}
