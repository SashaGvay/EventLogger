package service

import "event_logger/internal/domain/log"

type Service struct {
	Repository log.Repository
}

func New(repository log.Repository) *Service {
	return &Service{
		Repository: repository,
	}
}
