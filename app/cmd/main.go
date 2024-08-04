package main

import (
	"context"
	"log"

	_ "github.com/lib/pq"

	"github.com/Prrromanssss/chat-server/internal/app"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Println("Starting auth app")

	app, err := app.NewApp(ctx)
	if err != nil {
		log.Panicf("Cannot start app: %v", err)
	}

	if err = app.Run(ctx, cancel); err != nil {
		log.Panicf("Cannot start auth app: %v", err)
	}
}
