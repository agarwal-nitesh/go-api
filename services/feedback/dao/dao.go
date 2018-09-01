package dao

import (
	"gopkg.in/mgo.v2"
)

type FeedbackDao struct {
	FeedbackItem *FeedbackItem
	Feedback *Feedback
}

func NewFeedbackDao(mongo *mgo.Database) (*FeedbackDao, error) {
	return &FeedbackDao{
		&FeedbackItem{ Mongo: mongo},
		&Feedback{ Mongo: mongo},
	}, nil
}
