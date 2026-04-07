package model

import "time"

type Summary struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Events    int       `json:"events"`
	StartTime time.Time `json:"start_time"`
}
