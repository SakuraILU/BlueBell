package control

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code Code        `json:"Code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"Data"`
}

func WriteErrorResponse(ctx *gin.Context, code Code) {
	ctx.JSON(http.StatusBadRequest, ResponseData{
		Code: code,
		Msg:  code.string(),
		Data: nil,
	})
}

func WriteSuccessResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, ResponseData{
		Code: Success,
		Msg:  Success.string(),
		Data: data,
	})
}
