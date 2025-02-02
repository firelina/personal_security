package models

type UpdateRequest struct {
	EventID int    `json:"event_id"`
	Status  string `json:"status"`
}
