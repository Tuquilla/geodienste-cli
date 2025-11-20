package models

type Collection struct {
	Id string `json:"id"`
}

type Collections struct {
	Collections []Collection `json:"collections"`
}
