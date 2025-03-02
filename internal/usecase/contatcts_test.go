package usecase

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"personal_security/internal/domain"
	"testing"
)

type MockContactRepository struct {
	mock.Mock
}

func (m *MockContactRepository) CreateContact(ctx context.Context, newContact *domain.Contact) (int, error) {
	args := m.Called(ctx, newContact)
	return args.Int(0), args.Error(1)
}

func (m *MockContactRepository) GetContacts(ctx context.Context, userID int) ([]*domain.Contact, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]*domain.Contact), args.Error(1)
}

func TestCreateContact(t *testing.T) {
	mockRepo := new(MockContactRepository)
	contactService := NewContactService(mockRepo)
	t.Run("success create contact", func(t *testing.T) {
		newContact := &domain.Contact{UserID: 1, Name: "John Doe", Email: "john@example.com", Phone: "1234567890"}

		mockRepo.On("CreateContact", mock.Anything, newContact).Return(1, nil)

		contactID, err := contactService.CreateContact(context.Background(), newContact)

		assert.NoError(t, err)
		assert.Equal(t, 1, contactID)

		mockRepo.AssertExpectations(t)
	})
}

func TestGetContacts(t *testing.T) {
	mockRepo := new(MockContactRepository)
	contactService := NewContactService(mockRepo)

	t.Run("success get contacts", func(t *testing.T) {
		userID := 1
		expectedContacts := []*domain.Contact{
			{ID: 1, UserID: userID, Name: "John Doe", Email: "john@example.com", Phone: "1234567890"},
			{ID: 2, UserID: userID, Name: "Jane Doe", Email: "jane@example.com", Phone: "0987654321"},
		}

		mockRepo.On("GetContacts", mock.Anything, userID).Return(expectedContacts, nil)

		contacts, err := contactService.GetContacts(context.Background(), userID)

		assert.NoError(t, err)
		assert.Equal(t, expectedContacts, contacts)

		mockRepo.AssertExpectations(t)
	})
	t.Run("fail get contacts no contacts found", func(t *testing.T) {
		userID := 2

		mockRepo.On("GetContacts", mock.Anything, userID).Return([]*domain.Contact{}, errors.New("no contacts found"))

		contacts, err := contactService.GetContacts(context.Background(), userID)

		assert.Error(t, err)
		assert.Len(t, contacts, 0)
		assert.Equal(t, "no contacts found", err.Error())

		mockRepo.AssertExpectations(t)
	})
}
