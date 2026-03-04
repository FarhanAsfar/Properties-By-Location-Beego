package models

type Property struct {
	ID           string `json:"id"`
	Name         string `json:"Name"`
	PropertyName string `json:"property_name"`
	Country      string `json:"Country"`
	City         string `json:"City"`
	Slug         string `json:"slug"`
}
