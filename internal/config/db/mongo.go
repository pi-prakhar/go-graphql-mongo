package db

import (
	"context"
	"time"

	"github.com/pi-prakhar/go-graphql-mongo/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectionString string = "mongodb://localhost:27017"
var DBName = "graphql-job-board"
var DBCollection = "jobs"
var Client *mongo.Client

func Connect() {

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		logger.Log.Error("Error : Failed creating Mongo DB client", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		logger.Log.Error("Error : Failed Connecting to Mongo DB", err)
	}
	logger.Log.Info("Connected to Mongo DB")

	Client = client
}
