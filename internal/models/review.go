package models

type Review struct {
	Rating  int    `json:"rating"` // Rating 1 ~ 5
	Comment string `json:"comment"`
}
