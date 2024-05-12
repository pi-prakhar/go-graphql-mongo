package repository

import (
	"context"
	"time"

	"github.com/pi-prakhar/go-graphql-mongo/internal/config/db"
	model "github.com/pi-prakhar/go-graphql-mongo/internal/models"
	"github.com/pi-prakhar/go-graphql-mongo/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetJob(id string) (*model.JobListing, error) {
	jobCollec := db.Client.Database(db.DBName).Collection(db.DBCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	var jobListing model.JobListing
	err := jobCollec.FindOne(ctx, filter).Decode(&jobListing)
	if err != nil {
		logger.Log.Info("Error : Failed to fetch data from DB")
		return &jobListing, err
	}
	logger.Log.Info("Successfully fetched data from DB")
	return &jobListing, nil
}

func GetJobs() ([]*model.JobListing, error) {
	jobCollec := db.Client.Database(db.DBName).Collection(db.DBCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var jobListings []*model.JobListing
	cursor, err := jobCollec.Find(ctx, bson.D{})
	if err != nil {
		logger.Log.Info("Error : Failed to fetch data from DB")
		return jobListings, err
	}

	if err = cursor.All(context.TODO(), &jobListings); err != nil {
		logger.Log.Info("Error : Failed to parse DB data to models")
		return jobListings, err
	}
	logger.Log.Info("Successfully fetched data from DB")
	return jobListings, nil
}

func CreateJobListing(jobInfo model.CreateJobListingInput) (*model.JobListing, error) {
	var returnJobListing model.JobListing
	jobCollec := db.Client.Database(db.DBName).Collection(db.DBCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	inserg, err := jobCollec.InsertOne(ctx, bson.M{"title": jobInfo.Title, "description": jobInfo.Description, "url": jobInfo.URL, "company": jobInfo.Company})

	if err != nil {
		logger.Log.Info("Error : Failed to add data to DB")
		return &returnJobListing, err
	}

	insertedID := inserg.InsertedID.(primitive.ObjectID).Hex()
	returnJobListing = model.JobListing{ID: insertedID, Title: jobInfo.Title, Company: jobInfo.Company, Description: jobInfo.Description, URL: jobInfo.URL}
	logger.Log.Info("Successfully created data in DB")
	return &returnJobListing, nil
}

func UpdateJobListing(jobId string, jobInfo model.UpdateJobListingInput) (*model.JobListing, error) {
	jobCollec := db.Client.Database(db.DBName).Collection(db.DBCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	updateJobInfo := bson.M{}

	if jobInfo.Title != nil {
		updateJobInfo["title"] = jobInfo.Title
	}
	if jobInfo.Description != nil {
		updateJobInfo["description"] = jobInfo.Description
	}
	if jobInfo.URL != nil {
		updateJobInfo["url"] = jobInfo.URL
	}
	if jobInfo.Company != nil {
		updateJobInfo["company"] = jobInfo.Company
	}

	_id, _ := primitive.ObjectIDFromHex(jobId)
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": updateJobInfo}

	results := jobCollec.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var jobListing model.JobListing

	if err := results.Decode(&jobListing); err != nil {
		logger.Log.Info("Error : Failed to parse DB data to model")
		return &jobListing, err
	}
	logger.Log.Info("Successfully updated data in DB")
	return &jobListing, nil
}

func DeleteJobListing(jobId string) (*model.DeleteJobResponse, error) {
	jobCollec := db.Client.Database(db.DBName).Collection(db.DBCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(jobId)
	filter := bson.M{"_id": _id}
	_, err := jobCollec.DeleteOne(ctx, filter)
	if err != nil {
		logger.Log.Debug("Error : Failed to delete data from DB")
		return &model.DeleteJobResponse{DeletedJobID: jobId}, err
	}
	logger.Log.Info("Successfully deleted data from DB")
	return &model.DeleteJobResponse{DeletedJobID: jobId}, nil
}
