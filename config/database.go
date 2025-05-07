package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	MONGO_DB_URI  = "mongodb://localhost:47017"
	DATABASE_NAME = "edts-sharing"
)

var DB *mongo.Database

func InitDatabase() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	clientOptions := options.Client()
	clientOptions.ApplyURI(MONGO_DB_URI)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	DB = client.Database(DATABASE_NAME)
}
