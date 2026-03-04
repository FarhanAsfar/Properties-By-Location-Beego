package models

// response from /api/v1/get-properties/:location
type LocationResponse struct {
	GeoInfo LocationGeoInfo `json:"GeoInfo"`
}

// each location entry
type LocationGeoInfo struct {
	Slug string `json:"Slug"`
}
