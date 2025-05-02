package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client = DBInstance()

var dbName string // add global dbName variable

func DBInstance() *mongo.Client {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Error loading .env file:", err)
	}

	url := os.Getenv("MONGO_URI")
	dbName = os.Getenv("MONGO_DB") // read DB name from env

	if url == "" || dbName == "" {
		log.Fatal("MONGO_URI or MONGO_DB not set in .env")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal("Mongo connection failed:", err)
	}

	return client
}

func OpenCollection(conn *mongo.Client, coll string) *mongo.Collection {
	return conn.Database(dbName).Collection(coll)
}
