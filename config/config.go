package config

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// LogFileName is the name of the log file.
var LogFileName = "application.log"

// InitMongoClient initializes and returns a MongoDB client.
func InitMongoClient() (*mongo.Client, error) {
	// Create Logs directory if it doesn't exist
	logsDir := "Logs"
	if _, err := os.Stat(logsDir); os.IsNotExist(err) {
		err := os.Mkdir(logsDir, 0755)
		if err != nil {
			log.Fatalf("failed to create Logs directory: %v", err)
		}
	}

	// Open the log file for appending
	logFilePath := filepath.Join(logsDir, LogFileName)
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}
	defer logFile.Close()

	// Set the log output to the log file
	log.SetOutput(logFile)

	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("dotenv file not found")
	}

	// Get MongoDB connection URI from environment variable
	uri := os.Getenv("MONGODB_URL")
	if uri == "" {
		log.Fatal("you must set MONGODB_URI")
	}

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
	}

	return client, nil
}
