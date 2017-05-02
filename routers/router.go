package routers

import (
	"social-reviews/controllers"
	"github.com/astaxie/beego"
)

func init() {
  beego.Router("/", &controllers.MainController{})
	beego.Router("/signup", &controllers.UsersController{},"get:New;post:Create")
  beego.Router("/reviews", &controllers.ReviewsController{}, "get:Index;post:Create")
  beego.Router("/login", &controllers.SessionController{}, "get:New;post:Create")
  beego.Router("/logout", &controllers.SessionController{}, "get:Destroy")
}
