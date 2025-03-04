package config

import (
	"fmt"

	env "github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Conf struct {
	ENV string `env:"ENV" envDefault:"PROD"`

	KafkaHost         string `env:"KAFKA_HOST" envDefault:"localhost"`
	KafkaPort         string `env:"KAFKA_PORT" envDefault:"9092"`
	KafkaWeatherTopic string `env:"KAFKA_WEATHER_TOPIC" envDefault:"weather"`

	PostgresUser     string `env:"POSTGRES_USER" envDefault:"logger"`
	PostgresPassword string `env:"POSTGRES_PASSWORD" envDefault:"logger"`
	PostgresDB       string `env:"POSTGRES_DB" envDefault:"log"`
	PostgresPort     string `env:"POSTGRES_PORT" envDefault:"5432"`
	PostgresHost     string `env:"POSTGRES_HOST" envDefault:"localhost"`
}

func NewConfig() (*Conf, error) {
	cfg := &Conf{}

	godotenv.Load()
	if err := env.Parse(cfg, env.Options{}); err != nil {
		return nil, fmt.Errorf("env.Parse: %w", err)
	}
	return cfg, nil
}
