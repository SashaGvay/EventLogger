package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	domainLog "event_logger/internal/domain/log"
	"event_logger/internal/domain/log/models"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer struct {
	broker string
	topic  string

	logService domainLog.Service
}

func New(broker, topic string, logService domainLog.Service) *Consumer {
	return &Consumer{
		broker:     broker,
		topic:      topic,
		logService: logService,
	}
}

func (p *Consumer) Consume(ctx context.Context) error {
	kafkaConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": p.broker,
		"group.id":          p.topic + "_group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return fmt.Errorf("kafka.NewConsumer: %w", err)
	}
	defer kafkaConsumer.Close()

	err = kafkaConsumer.SubscribeTopics([]string{p.topic}, nil)
	if err != nil {
		return fmt.Errorf("kafkaConsumer.SubscribeTopics: %w", err)
	}

	for {
		select {
		case <-ctx.Done():
			log.Println(ctx.Err())
			return nil
		default:
			msg, err := kafkaConsumer.ReadMessage(-1)
			if err == nil {
				var weather models.Weather
				if err := json.Unmarshal(msg.Value, &weather); err != nil {
					log.Printf("json.Unmarshal: %v", err)
					continue
				}

				if err := p.logService.StoreWeather(ctx, &weather); err != nil {
					log.Printf("p.logService.StoreWeather: %v", err)
					continue
				}

				log.Println("Consumed msg")
			} else {
				log.Printf("kafkaConsumer.ReadMessage: %v", err)
			}
		}
	}
}
