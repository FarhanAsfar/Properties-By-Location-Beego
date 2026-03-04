package routers

import (
	"properties-by-location/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {

	ns := web.NewNamespace("/api/v1",
		web.NSRouter("/get-properties/:location", &controllers.PropertyController{}, "get:Get"),
	)
	web.AddNamespace(ns)
}
