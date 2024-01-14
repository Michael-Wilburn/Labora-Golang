package models

type Movie struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Created string  `json:"created"`
	Stock   int     `json:"stock"`
	Price   float64 `json:"price"`
}
