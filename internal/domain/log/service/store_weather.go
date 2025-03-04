package service

import (
	"context"
	"fmt"

	"event_logger/internal/domain/log/models"
)

func (s *Service) StoreWeather(ctx context.Context, weather *models.Weather) error {
	err := s.Repository.Put(ctx, weather)
	if err != nil {
		return fmt.Errorf("repository.Put: %w", err)
	}

	return nil
}
