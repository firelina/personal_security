package models

import "personal_security/internal/domain"

type ContactResponseItem struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Email  string `json:"email"`
}

type ContactResponse []ContactResponseItem

func NewContactResponse(contacts []*domain.Contact) ContactResponse {
	var response []ContactResponseItem
	for _, c := range contacts {
		response = append(response, ContactResponseItem{
			ID:     c.ID,
			UserID: c.UserID,
			Name:   c.Name,
			Phone:  c.Phone,
			Email:  c.Email,
		})
	}
	return response
}
