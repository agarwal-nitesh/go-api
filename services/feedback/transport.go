package feedback

import (
	"github.com/therudite/api/errors"
	"github.com/therudite/api/models/feedback"
	"github.com/therudite/api/services/feedback/utils"
	"github.com/gin-gonic/gin"
	"reflect"
)

type HTTPTransportInterface interface {
	// RequestMapper takes an interface as a param and binds the request to that interface
	DecodeRequest(requestType reflect.Type, ctx *gin.Context) (interface{}, error)

	// ResponseMapper takes an object and converts to API response
	EncodeResponse(res interface{}, err error) (interface{}, error)

	// MakeAPIHandler it returns http handler
	MakeAPIHandler() (func(ctx *gin.Context), error)
}

// Submit Feedback
// @Tags Admin V1
// Submit Feedback godoc
// @Summary Submit Feedback
// @Description Submit Feedback
// @Accept  json
// @Produce  json
// @Param feedback body models.Feedback true "Feedback"
// @Router /v1/feedback [PUT]
// @Success 200 {object} models.ResponseStruct
func (this *CreateSubmitFeedbackAPI) MakeAPIHandler() (func(ctx *gin.Context), error)  {
	return func(ctx *gin.Context) {
		req, err := this.DecodeRequest(reflect.TypeOf(models.Feedback{}), ctx)
		httputil := utils.HTTPUtil{ctx}
		if err != nil {
			httputil.HandleRequestMapperError(err)
			return
		}
		if request, ok := req.(*models.Feedback); ok {
			res, err := this.FeedbackServiceInterface.SubmitFeedback(request)
			httputil.HandleResponseMapper(this.EncodeResponse(res, err))
			return
		}
		panic(errors.RequestMapperTypeCastError.Error())
	}, nil
}

type CreateSubmitFeedbackAPI struct {
	DefaultHTTPTransport
}

// Get FeedbackItems
// @Tags Admin V1
// Get FeedbackItems godoc
// @Summary Get FeedbackItems
// @Description Get FeedbackItems
// @Accept  json
// @Produce  json
// @Router /v1/feedback [GET]
// @Success 200 {object} models.ResponseStruct
func (this *GetFeedbackOptionsAPI) MakeAPIHandler() (func(ctx *gin.Context), error)  {
	return func(ctx *gin.Context) {
		httputil := utils.HTTPUtil{ctx}
		res, err := this.FeedbackServiceInterface.GetFeedbackOptions()
		httputil.HandleResponseMapper(this.EncodeResponse(res, err))
		return
	}, nil
}

type GetFeedbackOptionsAPI struct {
	DefaultHTTPTransport
}