package controllers

// "encoding/json"
import (
  "fmt"
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
  email := c.GetString("email")
  username := c.GetString("username")

  user, err := models.CreateUser(username, email)
  if err != nil {
    c.Ctx.WriteString("failed to create user")
  }

  title := c.GetString("title")
  comment := c.GetString("comment")

  review, err := models.CreateReview(user.ID, title, comment)
  if err != nil {
    c.Ctx.WriteString("failed to create review")
  }

  fmt.Println("=============================")
  fmt.Println("created Review is:")
  fmt.Println(review.Comment)

  c.Data["json"] = map[string]interface{}{"review": review, "user": user}
  c.ServeJSON()
}
