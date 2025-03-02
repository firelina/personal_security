package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"personal_security/internal/domain"
)

type ReminderRepositoryInterface interface {
	CreateReminder(ctx context.Context, newReminder *domain.Reminder) (int, error)
	GetReminders(ctx context.Context, eventID int) ([]*domain.Reminder, error)
}

type ReminderRepository struct {
	pool *pgxpool.Pool
}

func NewReminderRepository(pool *pgxpool.Pool) *ReminderRepository {
	return &ReminderRepository{pool: pool}
}

func (r *ReminderRepository) CreateReminder(ctx context.Context, newReminder *domain.Reminder) (int, error) {
	query := `INSERT INTO personal_security.reminders (event_id, reminder_time, notification_method) VALUES ($1, $2, $3) RETURNING id`
	err := r.pool.QueryRow(ctx, query, newReminder.EventID, newReminder.ReminderTime, newReminder.NotificationMethod).Scan(&newReminder.ID)
	if err != nil {
		return 0, fmt.Errorf("failed to create reminder: %w", err)
	}
	return newReminder.ID, nil
}

func (r *ReminderRepository) GetReminders(ctx context.Context, eventID int) ([]*domain.Reminder, error) {
	query := `SELECT id, event_id, reminder_time, notification_method FROM personal_security.reminders WHERE event_id = $1`
	rows, err := r.pool.Query(ctx, query, eventID)
	if err != nil {
		return nil, fmt.Errorf("failed to get reminders: %w", err)
	}
	defer rows.Close()

	var reminders []*domain.Reminder
	for rows.Next() {
		reminder := &domain.Reminder{}
		if err := rows.Scan(&reminder.ID, &reminder.EventID, &reminder.ReminderTime, &reminder.NotificationMethod); err != nil {
			return nil, fmt.Errorf("failed to scan reminder: %w", err)
		}
		reminders = append(reminders, reminder)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during rows iteration: %w", err)
	}

	return reminders, nil
}
