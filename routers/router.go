package routers

import (
	"github.com/ahsaniqbal94/beego-getstarted/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
