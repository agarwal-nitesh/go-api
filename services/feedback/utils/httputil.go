package utils

import (
	"github.com/therudite/api/errors"
	"github.com/therudite/api/models/common"
	"github.com/gin-gonic/gin"
)

type HTTPUtil struct {
	Ctx *gin.Context
}

func (httputil HTTPUtil) HandleRequestMapperError(err error) {
	if err != nil {
		response := models.ResponseStruct{Message: "Failed", Desc: errors.BadRequestError.Error(), Data: nil, StatusCode: 400}
		httputil.Ctx.JSON(400, response)
	}
}

func (httputil HTTPUtil) HandleResponseMapper(res interface{}, err error) {
	if err != nil {
		response := models.ResponseStruct{Message: "Failed", Desc: "Request Failed", Data: nil, StatusCode: 503}
		if customErr, ok := err.(*errors.CustomError); ok {
			response.StatusCode = customErr.ErrorCode
		}
		httputil.Ctx.JSON(503, response)
		return
	}
	httputil.Ctx.JSON(200, res)
	return
}
