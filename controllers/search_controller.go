package controllers

import "github.com/beego/beego/v2/server/web"

type SearchRedirectController struct {
	web.Controller
}

func (c *SearchRedirectController) Get() {
	location := c.GetString("location")

	// redirect back to home page if location is empty
	if location == "" {
		c.Redirect("/", 302)
		return
	}

	// redirect to the property page
	c.Redirect("/all/"+location, 302)
}
