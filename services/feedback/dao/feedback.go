package dao

import (
	"github.com/therudite/api/models/feedback"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Feedback struct {
	Mongo *mgo.Database
}

func (this *Feedback) InsertFeedback(feedback *models.Feedback) (bool, error) {
	err := this.Mongo.C("feedback").Insert(feedback)

	if err != nil {
		return false, err
	}
	return true, nil
}

func (this *Feedback) FeedbackExists(id string, userId int64) (bool, error) {
	count, err := this.Mongo.C("feedback").Find(bson.M{"id": id, "user_id": userId}).Count()
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (this *Feedback) FeedbackUpdate(feedback *models.Feedback) (bool, error) {
	feedbackExists, err := this.FeedbackExists(feedback.Id, feedback.UserId)
	if feedbackExists {
		query := bson.M{"id": feedback.Id, "user_id": feedback.UserId}
		updatedBSON := bson.M{"$set": bson.M{"values": feedback.Values, "comments": feedback.Comments}}
		err = this.Mongo.C("feedback").Update(query, updatedBSON)
		if err != nil {
			return false, err
		}
		return true, nil
	} else {
		return this.InsertFeedback(feedback)
	}
}
