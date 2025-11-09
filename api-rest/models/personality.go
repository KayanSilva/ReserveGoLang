package models

type Personality struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	History string `json:"history"`
}
