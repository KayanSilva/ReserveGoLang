package models

import "fmt"

type Item struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

func (i Item) Info() string {
	return fmt.Sprintf("ID: %d, Name: %s, Quantity: %d, Price: %.2f", i.Id, i.Name, i.Quantity, i.Price)
}
