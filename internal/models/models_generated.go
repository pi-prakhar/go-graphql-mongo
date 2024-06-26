// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateJobListingInput struct {
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Company     string `json:"company" bson:"company"`
	URL         string `json:"url" bson:"url"`
}

type DeleteJobResponse struct {
	DeletedJobID string `json:"deletedJobId" bson:"deletedJobId"`
}

type JobListing struct {
	ID          string `json:"_id" bson:"_id"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Company     string `json:"company" bson:"company"`
	URL         string `json:"url" bson:"url"`
}

type Mutation struct {
}

type Query struct {
}

type UpdateJobListingInput struct {
	Title       *string `json:"title,omitempty" bson:"title"`
	Description *string `json:"description,omitempty" bson:"description"`
	Company     *string `json:"company,omitempty" bson:"company"`
	URL         *string `json:"url,omitempty" bson:"url"`
}
