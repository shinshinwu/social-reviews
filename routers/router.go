package routers

import (
	"social-reviews/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/reviews", &controllers.MainController{}, "get:Reviews")
}
