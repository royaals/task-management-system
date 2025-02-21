// internal/database/db.go
package database

import (
    "context"
    "log"
    "os"
    "time"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var (
    DB     *mongo.Database
    Client *mongo.Client
)

func InitDatabase() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Get MongoDB URI from environment variable
    mongoURI := os.Getenv("MONGODB_URI")
    if mongoURI == "" {
        log.Fatal("MONGODB_URI not set in environment")
    }

    // Connect to MongoDB
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
    if err != nil {
        log.Fatal("Failed to connect to MongoDB:", err)
    }

    // Ping the database
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal("Failed to ping MongoDB:", err)
    }

    // Set the database
    dbName := os.Getenv("DB_NAME")
    if dbName == "" {
        dbName = "task_management"
    }

    DB = client.Database(dbName)
    Client = client

    log.Println("Connected to MongoDB successfully")
}

// Get collection helper
func GetCollection(name string) *mongo.Collection {
    return DB.Collection(name)
}