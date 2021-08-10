package routers

import (
	"quickstart/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/phones", &controllers.PhonesController{})
	beego.Router("/download", &controllers.ExcelController{})
}
