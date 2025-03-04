package repository

import (
	"context"
	"fmt"
	"time"

	"event_logger/internal/domain/log/models"
)

func (r *Repository) Put(ctx context.Context, weather *models.Weather) error {
	query := `INSERT INTO weather_logs (temperature, timestamp) VALUES ($1, $2)`

	_, err := r.db.Exec(ctx, query, weather.Temperature, time.Now())
	if err != nil {
		return fmt.Errorf("db.Exec: %w", err)
	}

	return nil
}
