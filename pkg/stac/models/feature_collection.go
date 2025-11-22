package models

type FeatureCollection struct {
	Features []Feature `json:"features"`
}

type Feature struct {
	Assets map[string]AssetObject `json:"assets"`
}

type AssetObject struct {
	Href  string `json:"href"`
	Title string `json:"title"`
	Type  string `json:"type"`
}
