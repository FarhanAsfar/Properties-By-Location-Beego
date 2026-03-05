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

// carries the result through the channel of the location api call
type locationResult struct {
	slug string
	err  error
}

// carries outcome of category api call
type categoryResult struct {
	items []models.CategoryItem
	err   error
}

func GetPropertiesByLocation(location string) ([]models.Property, error) {

	//timeout context for both goroutines
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	baseURL, _ := web.AppConfig.String("travel_api_base_url")
	origin, _ := web.AppConfig.String("travel_api_origin")

	// separate goroutine for location api call
	locationChannel := make(chan locationResult, 1)
	go fetchLocation(ctx, baseURL, location, locationChannel)

	// block until we get the location goroutine response
	var slug string
	select {
	case res := <-locationChannel:
		if res.err != nil {
			return nil, res.err
		}
		slug = res.slug

	case <-ctx.Done():
		return nil, fmt.Errorf("Timed out , waiting fo location response")
	}

	// convert slug separator '/' with ':'
	categorySlug := strings.ReplaceAll(slug, "/", ":")

	categoryChannel := make(chan categoryResult, 1)
	go fetchCategory(ctx, baseURL, origin, categorySlug, categoryChannel)

	// block until the category goroutine responds
	select {
	case res := <-categoryChannel:
		if res.err != nil {
			return nil, res.err
		}
		return flattenProperties(res.items), nil
	case <-ctx.Done():
		return nil, fmt.Errorf("Timed out, waiting for properties response")
	}
}

func fetchLocation(ctx context.Context, baseURL string, keyword string, out chan<- locationResult) {
	url := fmt.Sprintf("%s/v1/location?keyword=%s", baseURL, keyword)

	resp, err := utils.FetchandDecode[models.LocationResponse](ctx, url, nil)
	if err != nil {
		out <- locationResult{err: fmt.Errorf("location API error: %w", err)}
		return
	}

	if resp == nil || resp.GeoInfo.LocationSlug == "" {
		out <- locationResult{err: fmt.Errorf("no location slug found for: %s", keyword)}
		return
	}

	out <- locationResult{slug: resp.GeoInfo.LocationSlug}
}

func fetchCategory(ctx context.Context, baseURL string, origin string, slug string, out chan<- categoryResult) {
	url := fmt.Sprintf("%s/v1/category/details/%s?items=1", baseURL, slug)

	headers := map[string]string{
		"Origin": origin,
	}

	resp, err := utils.FetchandDecode[models.CategoryResponse](ctx, url, headers)
	if err != nil {
		out <- categoryResult{err: fmt.Errorf("category API error: %w", err)}
		return
	}

	if resp == nil {
		out <- categoryResult{err: fmt.Errorf("empty response from category API")}
		return
	}

	out <- categoryResult{items: resp.Result.Items}
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
