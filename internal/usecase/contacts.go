package usecase

import (
	"personal_security/internal/domain"
)

type Contact struct {
}

var contacts []*domain.Contact

func (c *Contact) CreateContact(newContact *domain.Contact) int {
	newContact.ID = len(contacts) + 1
	contacts = append(contacts, newContact)
	return newContact.ID
}

func (c *Contact) GetContacts() []*domain.Contact {
	return contacts
}
