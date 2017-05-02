package controllers

import (
  "fmt"
  "strconv"
	"github.com/astaxie/beego"
  "social-reviews/models"
)

type ReviewsController struct {
	beego.Controller
}

func (c *ReviewsController) Prepare() {
  c.Layout = "layout/application.tpl"
}

func (c *ReviewsController) Index() {
  currentUser := c.GetSession("current_user")
  reviews, err := models.AllReviews()
  if err != nil {
    beego.BeeLogger.Error("Error finding reviews %v", err.Error())
  }
	c.Data["PageTitle"] = "Review Index Page"
  c.Data["Reviews"] = reviews
  c.Data["CurrentUser"] = currentUser
	c.TplName = "reviews.tpl"
}

func (c *ReviewsController) Create() {
  currentUser := c.GetSession("current_user")
  userId, _ := strconv.Atoi(currentUser.(struct{ID string}).ID)

  title := c.GetString("title")
  comment := c.GetString("comment")

  review, err := models.CreateReview(userId, title, comment)
  if err != nil {
    c.Ctx.WriteString("failed to create review")
  }

  fmt.Println("=============================")
  fmt.Println("created Review is:")
  fmt.Println(review.Comment)

  c.Data["json"] = map[string]interface{}{"review": review, "user": currentUser}
  c.ServeJSON()
}
