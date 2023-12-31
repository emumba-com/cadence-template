package utils

import (
	"github.com/gin-gonic/gin"
	"orderService/models"
)

type EmptyObj struct{}

func BuildResponse(ctx *gin.Context, statusCode int, status string, err string, data interface{}) {
	if err != "" {
		data = EmptyObj{}
	}

	response := models.Response{
		Status: status,
		Errors: err,
		Data:   data,
	}

	ctx.JSON(statusCode, response)
}

func BuildResponseAndAbort(ctx *gin.Context, statusCode int, status string, err string, data interface{}) {
	response := models.Response{
		Status: status,
		Errors: err,
		Data:   data,
	}

	ctx.AbortWithStatusJSON(statusCode, response)
}

func BuildDBErrResponse(ctx *gin.Context, err error, ph string) {
	statusCode, errMsg := ParseDBError(err, ph)
	BuildResponseAndAbort(ctx, statusCode, ERROR, errMsg, EmptyObj{})
}
