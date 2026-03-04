package models

// response from TravelAPI category endpoint
type CategoryResponse struct {
	Result CategoryResult `json:"Result"`
}

type CategoryResult struct {
	Count int            `json:"Count"`
	Items []CategoryItem `json:"Items"`
}

type CategoryItem struct {
	ID       string       `json:"ID"`
	GeoInfo  ItemGeoInfo  `json:"GeoInfo"`
	Property ItemProperty `json:"Property"`
}
type ItemGeoInfo struct {
	City    string `json:"City"`
	Country string `json:"Country"`
	Display string `json:"Display"`
}

type ItemProperty struct {
	PropertyName string  `json:"PropertyName"`
	PropertySlug string  `json:"PropertySlug"`
	PropertyType string  `json:"PropertyType"`
	Price        float64 `json:"Price"`
	ReviewScore  float64 `json:"ReviewScore"`
	StarRating   int     `json:"StarRating"`
	FeatureImage string  `json:"FeatureImage"`
}

type ItemPartner struct {
	URL string `json:"URL"`
}
