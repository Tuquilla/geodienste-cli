package models

type Collection struct {
	Id    string `json:"id"`
	Links []Link `json:"links"`
}

type Collections struct {
	Collections []Collection `json:"collections"`
}

type Link struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}
