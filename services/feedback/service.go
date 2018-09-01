package feedback

import (
	"github.com/therudite/api/errors"
	"github.com/therudite/api/resources"
	"github.com/therudite/api/models/feedback"
	"github.com/therudite/api/services/feedback/dao"
	"gopkg.in/mgo.v2"
	"time"
)

type FeedbackServiceInterface interface {
	GetFeedbackOptions() ([]models.FeedbackItem, error)
	SubmitFeedback(*models.Feedback) (bool, error)
}

type FeedbackService struct {
	Dao *dao.FeedbackDao
}

func NewService(resourcemanager resources.ResourceManagerInterface) (FeedbackServiceInterface, error) {
	var res interface{}
	var err error
	res, err = resourcemanager.Get("mongo")
	if err != nil {
		return nil, err
	}
	mongo, ok := res.(*mgo.Database)
	if !ok {
		return nil, errors.ResourceInitializationError
	}

	domain, err := dao.NewFeedbackDao(mongo)
	if err != nil {
		return nil, errors.DomainInitializationError
	}

	service := &FeedbackService{
		Dao: domain,
	}
	return service, nil
}

func (this *FeedbackService) GetFeedbackOptions() ([]models.FeedbackItem, error) {
	return this.Dao.FeedbackItem.GetFeedbackItems()
}

func (this *FeedbackService) SubmitFeedback(feedback *models.Feedback) (bool, error) {
	timeStamp := time.Now().Unix()
	feedbackItems, err := this.Dao.FeedbackItem.GetFeedbackItems()
	if err != nil {
		return false, err
	}
	var validatedFeedback = new(models.Feedback)
	for _, element := range feedback.Values {
		for _, item := range (feedbackItems) {
			if item.Id == element.Id {
				validatedFeedback.Values = append(validatedFeedback.Values,
					models.FeedbackItemResponse{Id: element.Id, Response: element.Response})
			}
		}
	}
	validatedFeedback.Id = feedback.Id
	validatedFeedback.Type = feedback.Type
	validatedFeedback.SubType = feedback.SubType
	validatedFeedback.LanguageCode = feedback.LanguageCode
	validatedFeedback.Comments = feedback.Comments
	validatedFeedback.UserId = feedback.UserId
	validatedFeedback.UserHandle = feedback.UserHandle
	validatedFeedback.Region = feedback.Region
	validatedFeedback.TimeStamp = timeStamp
	return this.Dao.Feedback.FeedbackUpdate(validatedFeedback)
}

