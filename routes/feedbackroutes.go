package routes

import (
	"github.com/therudite/api/resources"
	"github.com/therudite/api/services/feedback"
	"github.com/gin-gonic/gin"
)


func CreateFeedbackRoutes(engine *gin.Engine, resourceManager resources.ResourceManagerInterface) {
	var err error

	// instantiating services
	var feedbackService feedback.FeedbackServiceInterface
	feedbackService, err = feedback.NewService(resourceManager)

	if err != nil {
		panic("Error instantiating feedback service")
	}
	var transports feedback.TransportRoutes
	transports, err = feedback.MakeTransportRoutes(feedbackService)

	getFeedbackTransport, err := transports.GetFeedbackOptionsAPI.MakeAPIHandler()
	submitFeedbackTransport, err := transports.SubmitFeedbackAPI.MakeAPIHandler()

	if err != nil {
		panic("error instantiating transports: " + err.Error())
	}

	feedbackServiceGroup := engine.Group("/v1")
	feedbackServiceGroup.GET("/feedback", getFeedbackTransport)
	feedbackServiceGroup.POST("/feedback", submitFeedbackTransport)
}