package main

import (
	"context"
	"fmt"

	"log"

	"github.com/Prrromanssss/chat-server/config"
	"github.com/Prrromanssss/chat-server/internal/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Println("Starting server")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}
	log.Println("Config loaded")

	s := server.NewServer(
		cfg,
	)

	log.Printf(
		"Server config - Addres: %s",
		fmt.Sprintf("%s:%s", cfg.GRPC.Host, cfg.GRPC.Port),
	)

	if err = s.Run(ctx, cancel); err != nil {
		log.Fatalf("Cannot start server: %v", err)
	}
}
