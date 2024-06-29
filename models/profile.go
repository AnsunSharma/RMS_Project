package models

type Profile struct {
	ID                string `json:"id" bson:"_id,omitempty"`
	Applicant         User   `json:"applicant" bson:"applicant"`
	ResumeFileAddress string `json:"resume_file_address" bson:"resume_file_address"`
	Skills            string `json:"skills" bson:"skills"`
	Education         string `json:"education" bson:"education"`
	Experience        string `json:"experience" bson:"experience"`
	Name              string `json:"name" bson:"name"`
	Email             string `json:"email" bson:"email"`
	Phone             string `json:"phone" bson:"phone"`
}
