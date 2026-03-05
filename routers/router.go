package routers

import (
	"properties-by-location/controllers"
	"properties-by-location/middlewares"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	// ---page routes---
	// homepage
	web.Router("/", &controllers.HomeController{}, "get:Get")

	// search redirect
	web.Router("/search", &controllers.SearchRedirectController{}, "get:Get")

	// properties page
	web.Router("/all/:location", &controllers.PageController{}, "get:Get")

	ns := web.NewNamespace("/api/v1",
		web.NSBefore(middlewares.ValidateAPIKey),
		web.NSRouter("/get-properties/:location", &controllers.PropertyController{}, "get:Get"),
	)
	web.AddNamespace(ns)
}
