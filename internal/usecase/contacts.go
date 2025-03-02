package usecase

import (
	"context"
	"personal_security/internal/domain"
	"personal_security/internal/repository"
)

type ContactService struct {
	repo repository.ContactRepositoryInterface
}

func NewContactService(repo repository.ContactRepositoryInterface) *ContactService {
	return &ContactService{repo: repo}
}

func (s *ContactService) CreateContact(ctx context.Context, newContact *domain.Contact) (int, error) {
	resultChan := make(chan struct {
		ID  int
		Err error
	})

	go func() {
		userID, err := s.repo.CreateContact(ctx, newContact)
		resultChan <- struct {
			ID  int
			Err error
		}{ID: userID, Err: err}
	}()

	result := <-resultChan
	return result.ID, result.Err
}

func (s *ContactService) GetContacts(ctx context.Context, userID int) ([]*domain.Contact, error) {
	resultChan := make(chan struct {
		Contacts []*domain.Contact
		Err      error
	})

	go func() {
		contacts, err := s.repo.GetContacts(ctx, userID)
		resultChan <- struct {
			Contacts []*domain.Contact
			Err      error
		}{Contacts: contacts, Err: err}
	}()

	result := <-resultChan
	return result.Contacts, result.Err
}
