package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Prrromanssss/chat-server/config"
	"github.com/Prrromanssss/chat-server/internal/server"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DBName,
		cfg.Postgres.SSLMode,
	)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	s := server.NewServer(
		cfg,
		db,
	)

	log.Printf(
		"Server config - Addres: %s",
		fmt.Sprintf("%s:%s", cfg.GRPC.Host, cfg.GRPC.Port),
	)

	if err = s.Run(ctx, cancel); err != nil {
		log.Fatalf("Cannot start server: %v", err)
	}
}
