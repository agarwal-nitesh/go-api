package dao

import (
	"github.com/therudite/api/models/feedback"
	"gopkg.in/mgo.v2"
)

type FeedbackItem struct {
	Mongo *mgo.Database
}

func (this *FeedbackItem) GetFeedbackItems() ([]models.FeedbackItem, error) {
	result := []models.FeedbackItem{}
	err := this.Mongo.C("feedback_items").Find(nil).Sort("priority").All(&result)

	if err != nil {
		return nil, err
	}
	return result, nil
}
