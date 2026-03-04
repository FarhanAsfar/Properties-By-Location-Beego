package controllers

import (
	"properties-by-location/services"
	"properties-by-location/utils"

	"github.com/beego/beego/v2/server/web"
)

// GET api/v1/get-properties/:location
type PropertyController struct {
	web.Controller
}

func (c *PropertyController) Get() {
	// get the location from the param
	location := c.Ctx.Input.Param(":location")

	if location == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = utils.ErrorResponse("Location is required")

		c.ServeJSON()
		return
	}

	properties, err := services.GetPropertiesByLocation(location)

	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = utils.ErrorResponse(err.Error())
		c.ServeJSON()

		return
	}

	// return the flattened list as json
	c.Data["json"] = utils.SuccessResponse(properties)
	c.ServeJSON()
}
