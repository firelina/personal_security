package models

import "time"

type CreateEventRequest struct {
	UserID      int       `json:"user_id"`
	Title       string    `json:"title"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
}
