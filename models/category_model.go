package models

// response from TravelAPI category endpoint
type CategoryResponse struct {
	Data CategoryData `json:"data"`
}

type CategoryData struct {
	Properties []PropertyItem `json:"properties"`
}

// each property returned by the api

type PropertyItem struct {
	ID           string `json:"id"`
	Name         string `json:"Name"`
	PropertyName string `json:"property_name"`
	Country      string `json:"Country"`
	City         string `json:"City"`
	Slug         string `json:"Slug"`
}
