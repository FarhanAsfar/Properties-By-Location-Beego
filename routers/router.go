package routers

import (
	"properties-by-location/controllers"
	"properties-by-location/middlewares"

	"github.com/beego/beego/v2/server/web"
)

func init() {

	ns := web.NewNamespace("/api/v1",
		web.NSBefore(middlewares.ValidateAPIKey),
		web.NSRouter("/get-properties/:location", &controllers.PropertyController{}, "get:Get"),
	)
	web.AddNamespace(ns)
}
