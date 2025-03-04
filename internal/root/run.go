package root

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func (r *Root) Run(ctx context.Context) error {
	log.Println("Starting...")
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		if r.producerWeather == nil {
			return
		}

		err := r.producerWeather.Produce(ctx)
		if err != nil {
			log.Fatalf("producerWeather.Produce error: %v", err)
		}
	}()

	go func() {
		if r.consumerWeather == nil {
			return
		}

		err := r.consumerWeather.Consume(ctx)
		if err != nil {
			log.Fatalf("r.consumerWeather.Consume error: %v", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	<-sigChan
	log.Println("Shutting down...")

	cancel()

	return nil
}
