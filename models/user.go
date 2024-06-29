package models

type User struct {
	ID              string `json:"id" bson:"_id,omitempty"`
	Name            string `json:"name" bson:"name"`
	Email           string `json:"email" bson:"email"`
	Password        string `json:"password" bson:"password"`
	UserType        string `json:"userType" bson:"userType"`
	ProfileHeadline string `json:"profile_headline" bson:"profile_headline"`
	Address         string `json:"address" bson:"address"`
	// Profile:        Profile `json:"profile_headline" bson:"profile_headline"`
}
