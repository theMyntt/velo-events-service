package database

import (
	"database/sql"
	"errors"

	"gitlab.com/velo-company/services/events-service/internal/core/ports"
)

type cancelSubscriptionAdapter struct {
	DB *sql.DB
}

func NewCancelSubscriptionAdapter(db *sql.DB) ports.CancelSubscriptionPort {
	return &cancelSubscriptionAdapter{
		DB: db,
	}
}

const (
	searchEventQuery = `SELECT 1 FROM tb_user_events WHERE fk_id_event = $1`
	cancelEventQuery = `UPDATE tb_user_events SET participation_status_event = $1 WHERE fk_id_event = $2`
)

func (c cancelSubscriptionAdapter) Execute(eventId int) error {
	var eventExists int
	err := c.DB.QueryRow(searchEventQuery, eventId).Scan(&eventExists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("esse evento não existe")
		}
		return err
	}

	tx, err := c.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.Exec(cancelEventQuery, 2, eventId)

	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}
