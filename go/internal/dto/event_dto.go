package dto

import "time"

type CreateEvent struct {
	UserId   int            `json:"user_id" binding:"required"`
	Action   string         `json:"action" binding:"required"`
	Metadata map[string]any `json:"metadata,omitempty"`
}

type FilterEvent struct {
	UserId    *int       `form:"user_id"`
	StartTime *time.Time `form:"start_time"`
	EndTime   *time.Time `form:"end_time"`
}
