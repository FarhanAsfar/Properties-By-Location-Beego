package main

import (
	"fmt"
	_ "properties-by-location/routers"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	// validating config keys
	mustConfig("travel_api_base_url")
	mustConfig("travel_api_origin")
	mustConfig("x_api_key")

	web.Run()
}

func mustConfig(key string) {
	val, err := web.AppConfig.String(key)

	if err != nil || val == "" {
		panic(fmt.Sprintf("Missing required config key: %s. Please add it into your app.conf file", key))
	}
}
