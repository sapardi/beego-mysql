package routers

import (
	"github.com/sapardi/beego-mysql/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
