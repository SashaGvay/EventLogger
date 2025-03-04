package log

import (
	"context"

	"event_logger/internal/domain/log/models"
)

type Service interface {
	StoreWeather(ctx context.Context, weather *models.Weather) error
}
