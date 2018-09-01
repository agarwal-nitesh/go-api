package feedback

import (
commonModels "github.com/therudite/api/models/common"
"github.com/gin-gonic/gin"
"reflect"
)

type DefaultHTTPTransport struct {
	FeedbackServiceInterface
}

func (this *DefaultHTTPTransport) DecodeRequest(requestType reflect.Type, ctx *gin.Context) (interface{}, error) {
	request := reflect.New(requestType).Interface()
	if err := ctx.BindJSON(&request); err == nil {
		return request, nil
	} else {
		return nil, err
	}
}

func (this *DefaultHTTPTransport) EncodeResponse(res interface{}, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}
	response := commonModels.ResponseStruct{Message: "Success", Desc: "Request Success", Data: res, StatusCode: 200}
	return response, nil
}