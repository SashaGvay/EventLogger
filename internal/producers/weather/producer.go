package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"event_logger/internal/domain/log/models"
	rate_limiter "github.com/SashaGvay/RateLimiter"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Producer struct {
	broker string
	topic  string
}

func New(broker, topic string) *Producer {
	return &Producer{
		broker: broker,
		topic:  topic,
	}
}

func (p *Producer) Produce(ctx context.Context) error {
	kafkaProducer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": p.broker,
	})
	if err != nil {
		return fmt.Errorf("kafka.NewProducer: %w", err)
	}

	limiter := rate_limiter.NewRateLimiter(5 * time.Second)

	go func() {
		defer limiter.Stop()

		for {
			err = limiter.Wait(ctx)
			if err != nil {
				log.Printf("limiter.Wait: %v", err)
				return
			}

			log.Println("Producing msg")

			weatherDto := &models.Weather{
				Temperature: math.Round((rand.Float64()*40-10)*10) / 10, // mocked temperature
			}

			payload, err := json.Marshal(weatherDto)
			if err != nil {
				return
			}

			err = kafkaProducer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &p.topic, Partition: kafka.PartitionAny},
				Value:          payload,
			}, nil)
			if err != nil {
				log.Println("kafkaProducer.Produce: ", err)
				return
			}
		}
	}()

	return nil
}
