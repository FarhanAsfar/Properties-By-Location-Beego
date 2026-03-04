package services

import (
	"context"
	"fmt"
	"properties-by-location/models"
	"properties-by-location/utils"
	"strings"
	"time"

	"github.com/beego/beego/v2/server/web"
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

// perform two external api calls
func fetchProperties(ctx context.Context, location string) ([]models.Property, error) {
	baseURL, _ := web.AppConfig.String("travel_api_base_url")
	fmt.Printf("Debug: url: [%s]\n", baseURL)
	origin, _ := web.AppConfig.String("travel_api_origin")

	// resolve location slug
	locationURL := fmt.Sprintf("%s/v1/location?keyword=%s", baseURL, location)

	locationResp, err := utils.FetchandDecode[models.LocationResponse](ctx, locationURL, nil)

	if err != nil {
		return nil, fmt.Errorf("Failed to fetch location: %w", err)
	}

	if locationResp == nil || len(locationResp.GeoInfo.Slug) == 0 {
		return nil, fmt.Errorf("No location data found for: %s", location)
	}

	// extract slug from the result
	rawSlug := locationResp.GeoInfo.Slug

	//replace '/' with ':'
	categorySlug := strings.ReplaceAll(rawSlug, "/", ":")

	categoryURL := fmt.Sprintf("%s/v1/category/details/%s?items=1", baseURL, categorySlug)

	headers := map[string]string{
		"Origin": origin,
	}

	categoryResp, err := utils.FetchandDecode[models.CategoryResponse](ctx, categoryURL, headers)

	if err != nil {
		return nil, fmt.Errorf("Failed to fetch properties: %w", err)
	}

	if categoryResp == nil {
		return nil, fmt.Errorf("Empty response from category API")
	}

	properties := flattenProperties(categoryResp.Result.Items)

	return properties, nil
}

// map PropertyItem structs to Property model
func flattenProperties(items []models.CategoryItem) []models.Property {
	result := make([]models.Property, 0, len(items))

	for _, item := range items {
		result = append(result, models.Property{
			ID:           item.ID,
			PropertyName: item.Property.PropertyName,
			PropertyType: item.Property.PropertyType,
			City:         item.GeoInfo.City,
			Country:      item.GeoInfo.Country,
			Price:        item.Property.Price,
			ReviewScore:  item.Property.ReviewScore,
			StarRating:   item.Property.StarRating,
			Slug:         item.Property.PropertySlug,
		})
	}
	return result
}
