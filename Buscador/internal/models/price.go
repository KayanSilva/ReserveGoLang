package models

import "time"

type Price struct {
	StoreName string    `json:"storeName"`
	Value     float64   `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
