package controllers

import (
	"context"
	"fmt"
	"properties-by-location/models"
	"properties-by-location/utils"
	"time"

	"github.com/beego/beego/v2/server/web"
)

// pageResult: would pass the internal api result through a channel
type pageResult struct {
	properties []models.Property
	err        error
}

// PageController: handles page route
// route: GET /all/:location
type PageController struct {
	web.Controller
}

// Get() calls the internal api and passes the result to the template
func (c *PageController) Get() {
	location := c.Ctx.Input.Param(":location")

	// validate location param
	if location == "" {
		c.Data["Error"] = "Location is required"
		c.TplName = "properties.tpl"

		return
	}

	//build the internal api url
	apiKey, _ := web.AppConfig.String("x_api_key")
	port, _ := web.AppConfig.String("httpport")

	internalURL := fmt.Sprintf("http://127.0.0.1:%s/api/v1/get-properties/%s", port, location)

	// channel to receive the result from goroutine
	resultChannel := make(chan pageResult, 1)

	// timeout for the internal http call
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// calling the api asynchronously
	go func() {
		props, err := callInternalAPI(ctx, internalURL, apiKey)
		resultChannel <- pageResult{properties: props, err: err}
	}()

	// waiting for result / timepout
	var props []models.Property
	var fetchErr error

	select {
	case res := <-resultChannel:
		props = res.properties
		fetchErr = res.err
	case <-ctx.Done():
		fetchErr = fmt.Errorf("Request time out")
	}

	// passing data to the template
	c.Data["Location"] = location

	if fetchErr != nil {
		c.Data["Error"] = fmt.Sprintf("Could not load properties: %s", fetchErr.Error())
	} else {
		c.Data["Properties"] = props
		c.Data["HasProperties"] = len(props) > 0
	}

	c.TplName = "properties.tpl"
}

// make an http get request to our internal endpoint
func callInternalAPI(ctx context.Context, url string, apiKey string) ([]models.Property, error) {
	headers := map[string]string{
		"X-Api-Key": apiKey,
	}

	resp, err := utils.FetchandDecode[utils.APIResponse[[]models.Property]](ctx, url, headers)

	if err != nil {
		return nil, err
	}

	// check if there's any error in response body
	if resp == nil || !resp.Success {
		msg := "Internal API error"

		if resp != nil && resp.Error != "" {
			msg = resp.Error
		}
		return nil, fmt.Errorf(msg)
	}

	if resp.Data == nil {
		return nil, fmt.Errorf("No data returned from the internal API")
	}

	return *resp.Data, nil
}

// HomeController: handles the homepage and search form
// route: GET /
type HomeController struct {
	web.Controller
}

func (c *HomeController) Get() {
	c.TplName = "index.tpl"
}
