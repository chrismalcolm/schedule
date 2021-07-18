package models

type Bus struct {
	ID       string   `json:"id"`
	Location Location `json:"location"`
	Metadata Metadata `json:"metadata"`
}

type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type Metadata struct {
	Manufacturer string `json:"manufacturer"`
}
