package main

import (
	_ "properties-by-location/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

