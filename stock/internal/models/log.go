package models

import "time"

type Log struct {
	Timestamp time.Time `json:"timestamp"`
	Action    string    `json:"action"`
	User      string    `json:"user"`
	ItemID    int       `json:"itemId"`
	Quantity  int       `json:"quantity"`
	Reason    string    `json:"reason"`
}
