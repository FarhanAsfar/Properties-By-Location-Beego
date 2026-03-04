package models

type Property struct {
	ID           string  `json:"id"`
	PropertyName string  `json:"property_name"`
	PropertyType string  `json:"property_type"`
	Country      string  `json:"country"`
	City         string  `json:"city"`
	Slug         string  `json:"slug"`
	URL          string  `json:"url"`
	StarRating   int     `json:"star_rating"`
	ReviewScore  float64 `json:"review_score"`
	Price        float64 `json:"price"`
}
