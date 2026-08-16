package root

import (
	"context"
	"fmt"

	pgx "github.com/jackc/pgx/v5"

	"event_logger/config"

	domainLog "event_logger/internal/domain/log"
	logRepository "event_logger/internal/domain/log/repository"
	logService "event_logger/internal/domain/log/service"

	consumerWeather "event_logger/internal/consumers/weather"
	producerWeather "event_logger/internal/producers/weather"
)

type Root struct {
	Cfg *config.Conf

	Infrastructure struct {
		DbInstance *pgx.Conn
	}

	Entity struct {
		Log struct {
			Service    domainLog.Service
			Repository domainLog.Repository
		}
	}

	producerWeather *producerWeather.Producer
	consumerWeather *consumerWeather.Consumer
}

func New(ctx context.Context, cfg *config.Conf) (*Root, error) {
	r := &Root{
		Cfg: cfg,
	}

	err := r.initInfrastructure(ctx)
	if err != nil {
		return nil, fmt.Errorf("r.initInfrastructure: %w", err)
	}

	err = r.initEntities(ctx)
	if err != nil {
		return nil, fmt.Errorf("r.initEntities: %w", err)
	}

	r.initProducer(ctx)
	r.initConsumer(ctx)

	return r, nil
}

func (r *Root) initInfrastructure(ctx context.Context) error {
	var err error

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		r.Cfg.PostgresUser,
		r.Cfg.PostgresPassword,
		r.Cfg.PostgresHost,
		r.Cfg.PostgresPort,
		r.Cfg.PostgresDB,
	)

	r.Infrastructure.DbInstance, err = pgx.Connect(ctx, dsn)
	if err != nil {
		return fmt.Errorf("pgx.Connect: %w", err)
	}

	err = r.Infrastructure.DbInstance.Ping(ctx)
	if err != nil {
		return fmt.Errorf("r.Infrastructure.DbInstance.Ping: %w", err)
	}

	return nil
}

func (r *Root) initEntities(_ context.Context) error {
	r.Entity.Log.Repository = logRepository.New(r.Infrastructure.DbInstance)
	r.Entity.Log.Service = logService.New(r.Entity.Log.Repository)

	return nil
}

func (r *Root) initProducer(_ context.Context) {
	r.producerWeather = producerWeather.New(
		fmt.Sprintf("%s:%s", r.Cfg.KafkaHost, r.Cfg.KafkaPort),
		r.Cfg.KafkaWeatherTopic,
	)
}
func (r *Root) initConsumer(_ context.Context) {
	r.consumerWeather = consumerWeather.New(
		fmt.Sprintf("%s:%s", r.Cfg.KafkaHost, r.Cfg.KafkaPort),
		r.Cfg.KafkaWeatherTopic,
		r.Entity.Log.Service,
	)
}
