package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"personal_security/internal/domain"
)

type EventRepositoryInterface interface {
	CreateEvent(ctx context.Context, newEvent *domain.Event) (int, error)
	GetEvents(ctx context.Context, userID int) ([]*domain.Event, error)
	UpdateEventStatus(ctx context.Context, eventID int, status string) (*domain.Event, error)
}

type EventRepository struct {
	pool *pgxpool.Pool
}

func NewEventRepository(pool *pgxpool.Pool) *EventRepository {
	return &EventRepository{pool: pool}
}

func (r *EventRepository) CreateEvent(ctx context.Context, newEvent *domain.Event) (int, error) {
	query := `INSERT INTO personal_security.events (user_id, title, date, description, status) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := r.pool.QueryRow(ctx, query, newEvent.UserID, newEvent.Title, newEvent.Date, newEvent.Description, newEvent.Status).Scan(&newEvent.ID)
	if err != nil {
		return 0, fmt.Errorf("failed to create event: %w", err)
	}
	return newEvent.ID, nil
}

func (r *EventRepository) GetEvents(ctx context.Context, userID int) ([]*domain.Event, error) {
	query := `SELECT id, user_id, title, date, description, status FROM personal_security.events WHERE user_id = $1`
	rows, err := r.pool.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get events: %w", err)
	}
	defer rows.Close()

	var events []*domain.Event
	for rows.Next() {
		event := &domain.Event{}
		if err := rows.Scan(&event.ID, &event.UserID, &event.Title, &event.Date, &event.Description, &event.Status); err != nil {
			return nil, fmt.Errorf("failed to scan event: %w", err)
		}
		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during rows iteration: %w", err)
	}

	return events, nil
}

func (r *EventRepository) UpdateEventStatus(ctx context.Context, eventID int, status string) (*domain.Event, error) {
	query := `UPDATE personal_security.events SET status = $1 WHERE id = $2 RETURNING id, user_id, title, date, description, status`
	event := &domain.Event{}
	err := r.pool.QueryRow(ctx, query, status, eventID).Scan(&event.ID, &event.UserID, &event.Title, &event.Date, &event.Description, &event.Status)
	if err != nil {
		return nil, fmt.Errorf("failed to update event status: %w", err)
	}
	return event, nil
}
