package services

import (
	"context"
	"fmt"
	"properties-by-location/models"
	"time"
)

type propertyResult struct {
	properties []models.Property
	err        error
}

func GetPropertiesByLocation(location string) ([]models.Property, error) {
	// channel for receiving result
	resultChannel := make(chan propertyResult, 1)

	// timeout for the fetch operation
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// asynchronous api call using goroutine
	go func() {
		props, err := fetchProperties(ctx, location)
		resultChannel <- propertyResult{properties: props, err: err}
	}()

	// waiting for the goroutine to return result
	select {
	case res := <-resultChannel:
		return res.properties, res.err
	case <-ctx.Done():
		return nil, fmt.Errorf("Request timed out while fetching properties")
	}
}
