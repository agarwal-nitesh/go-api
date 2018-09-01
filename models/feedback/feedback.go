package models

type Feedback struct {
	Id string `bson:"id" json:"id"`
	Type string `bson:"type" json:"type"`
	SubType string `bson:"sub_type" json:"sub_type"`
	LanguageCode string `bson:"language_code" json:"language_code"`
	UserId int64 `bson:"user_id" json:"user_id"`
	UserHandle string `bson:"user_handle" json:"user_handle"`
	Region string `bson:"region" json:"region"`
	TimeStamp int64 `bson:"time_stamp" json:"time_stamp"`
	Values []FeedbackItemResponse `bson:"values" json:"values"`
	Comments string `bson:"comments" json:"comments"`
}

type FeedbackItem struct {
	Id string `bson:"id" json:"id"`
	Name string `bson:"name" json:"name"`
	DataType string `bson:"data_type" json:"-"`
	WidgetType string `bson:"widget_type" json:"widget_type"`
}

type FeedbackItemResponse struct {
	Id string `bson:"id" json:"id"`
	Response int `bson:"response" json:"response"`
}

type FeedbackResponse struct {
	Title string `json:"title"`
	Subtitle string `json:"subtitle"`
	Items []FeedbackItem `json:"items"`
	CommentsEnabled bool `json:"comments_enabled"`
}

type FeedbackRequest struct {
	Id string `json:"id"`
	Type string `json:"type"`
	SubType string `json:"sub_type"`
	LanguageCode string `json:"language_code"`
	Values []FeedbackItemResponse `json:"values"`
	Comments string `json:"comments"`
}
