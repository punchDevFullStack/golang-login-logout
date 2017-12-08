package routers

import (
	"github.com/thanamat/golang-login-logout/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{}, "get:Get;post:Login")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")
}
