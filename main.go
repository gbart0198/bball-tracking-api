package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gbart0198/bball-tracker-api/api"
	"github.com/gbart0198/bball-tracker-api/storage"
	"github.com/joho/godotenv"
)

func main() {
	listenAddr := flag.String("listenAddr", ":8080", "server listen address")
	flag.Parse()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	if dbHost == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		log.Fatal("Missing required environment variables.")
	}

	connStr := fmt.Sprintf("Data Source=%s;Database=%s;Integrated Security=false;User ID=%s;Password=%s;", dbHost, dbName, dbUser, dbPassword)

	storage := storage.NewMSSQLStorage(connStr)

	fmt.Println("Listening on port ", *listenAddr)
	server := api.NewServer(*listenAddr, storage)
	storage.Connect()
	log.Fatal(server.Start())
}
