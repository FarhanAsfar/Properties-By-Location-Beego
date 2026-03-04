package models

// response from /api/v1/get-properties/:location
type LocationResponse struct {
	Data []LocationItem `json:"data"`
}

// each location entry
type LocationItem struct {
	Slug string `json:"LocationSlug"`
}
