package db

import (
	"context"
	"time"

	"github.com/pi-prakhar/go-graphql-mongo/pkg/logger"
	"github.com/pi-prakhar/utils/loader"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DBName = "graphql-job-board"
var DBCollection = "jobs"
var Client *mongo.Client

func getConnectionString() string {
	var connectionString string
	isDocker, err := loader.GetValueFromConf("docker")
	if err != nil {
		logger.Log.Error("Error : Failed to fetch docker property from configurations", err)
	}

	if isDocker == "true" {
		connectionString, err = loader.GetValueFromConf("docker-mongo-uri")
		if err != nil {
			logger.Log.Error("Error : Failed to fetch docker-mongo-uri property from configurations", err)
		}
	} else {
		connectionString, err = loader.GetValueFromConf("local-mongo-uri")
		if err != nil {
			logger.Log.Error("Error : Failed to fetch local-mongo-uri property from configurations", err)
		}
	}
	return connectionString
}

func Connect() {

	client, err := mongo.NewClient(options.Client().ApplyURI(getConnectionString()))
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
