package models

type CreateContactRequest struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Email  string `json:"email"`
}
