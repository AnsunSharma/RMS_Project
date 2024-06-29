package models

import "time"

type Job struct {
	ID                string    `json:"id" bson:"_id,omitempty"`
	Title             string    `json:"title" bson:"title"`
	Description       string    `json:"description" bson:"description"`
	PostedOn          time.Time `json:"posted_on" bson:"posted_on"`
	TotalApplications int       `json:"total_applications" bson:"total_applications"`
	CompanyName       string    `json:"company_name" bson:"company_name"`
	PostedBy          User      `json:"posted_by" bson:"posted_by"`
	CreatorName       string    `json:"CreatorName" bson:"CreatorName"`
}
