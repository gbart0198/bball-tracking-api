package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gbart0198/bball-tracker-api/api"
	"github.com/gbart0198/bball-tracker-api/storage"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	listenAddr := flag.String("listenAddr", ":8080", "server listen address")
	flag.Parse()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	if dbHost == "" || dbUser == "" || dbPassword == "" || dbName == "" || dbPort == "" {
		log.Fatal("Missing required environment variables.")
	}

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)

	storage := storage.NewPostgreSqlStorage(ctx, connStr)
	defer storage.Close()

	fmt.Println("Listening on port ", *listenAddr)
	server := api.NewServer(*listenAddr, storage)
	log.Fatal(server.Start())
}
