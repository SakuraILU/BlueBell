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

func WriteErrorResponse(c *gin.Context, code Code) {
	c.JSON(http.StatusBadRequest, ResponseData{
		Code: code,
		Msg:  code.string(),
		Data: nil,
	})
}

func WriteSuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code: Success,
		Msg:  Success.string(),
		Data: data,
	})
}
