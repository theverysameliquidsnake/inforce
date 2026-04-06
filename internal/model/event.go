package model

import "time"

type Event struct {
	Id        int            `json:"id"`
	UserId    int            `json:"user_id"`
	Action    string         `json:"action"`
	Timestamp time.Time      `json:"timestamp"`
	Metadata  map[string]any `json:"metadata,omitempty"`
}
