package log

import (
	"context"

	"event_logger/internal/domain/log/models"
)

type Repository interface {
	Put(ctx context.Context, weather *models.Weather) error
}
