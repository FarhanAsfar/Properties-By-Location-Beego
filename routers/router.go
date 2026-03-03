package routers

import (
	"properties-by-location/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/", &controllers.MainController{})
}
