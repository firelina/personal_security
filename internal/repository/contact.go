package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"personal_security/internal/domain"
)

type ContactRepositoryInterface interface {
	CreateContact(ctx context.Context, newContact *domain.Contact) (int, error)
	GetContacts(ctx context.Context, userID int) ([]*domain.Contact, error)
}

type ContactRepository struct {
	pool *pgxpool.Pool
}

func NewContactRepository(pool *pgxpool.Pool) *ContactRepository {
	return &ContactRepository{pool: pool}
}

func (r *ContactRepository) CreateContact(ctx context.Context, newContact *domain.Contact) (int, error) {
	query := `INSERT INTO personal_security.contacts (user_id, name, email, phone) VALUES ($1, $2, $3, $4) RETURNING id`
	err := r.pool.QueryRow(ctx, query, newContact.UserID, newContact.Name, newContact.Email, newContact.Phone).Scan(&newContact.ID)
	if err != nil {
		return 0, fmt.Errorf("failed to create contact: %w", err)
	}
	return newContact.ID, nil
}

func (r *ContactRepository) GetContacts(ctx context.Context, userID int) ([]*domain.Contact, error) {
	query := `SELECT id, user_id, name, email, phone FROM personal_security.contacts WHERE user_id = $1`
	rows, err := r.pool.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get contacts: %w", err)
	}
	defer rows.Close()

	var contacts []*domain.Contact
	for rows.Next() {
		contact := &domain.Contact{}
		if err := rows.Scan(&contact.ID, &contact.UserID, &contact.Name, &contact.Email, &contact.Phone); err != nil {
			return nil, fmt.Errorf("failed to scan contact: %w", err)
		}
		contacts = append(contacts, contact)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during rows iteration: %w", err)
	}

	return contacts, nil
}
