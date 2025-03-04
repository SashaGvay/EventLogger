package main

import (
	"context"
	"log"
	"os"

	"event_logger/config"
	"event_logger/internal/root"
)

func main() {
	ctx := context.Background()

	conf, err := config.NewConfig()
	if err != nil {
		log.Panicf("config.NewConfig: %v", err)
	}

	app, err := root.New(ctx, conf)
	if err != nil {
		log.Panicf("root.New: %v", err)
	}

	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		if err := app.RunMigrations(); err != nil {
			log.Panicf("app.RunMigrations: %v", err)
		}
		return
	}

	if err = app.Run(ctx); err != nil {
		log.Panicf("app.Run: %v", err)
	}
}
