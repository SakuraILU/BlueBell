package control

import (
	log "bluebell/Log"
	logic "bluebell/Logic"
	model "bluebell/Model"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	param := new(model.ParamSignUp)

	if err := c.ShouldBindJSON(param); err != nil {
		log.Error(err)
		WriteErrorResponse(c, InvalidParam)
		return
	}

	if !validateSignUp(param) {
		WriteErrorResponse(c, InvalidParam)
		return
	}

	log.Infof("try to regist user %v", param)

	// signup the user
	if err := logic.SignUp(param); err != nil {
		WriteErrorResponse(c, UserExist)
	}

	WriteSuccessResponse(c, nil)
}

func Login(c *gin.Context) {
	param := new(model.ParamLogin)

	if err := c.ShouldBindJSON(param); err != nil {
		log.Error(err)
		WriteErrorResponse(c, InvalidParam)
		return
	}

	if !validateLogin(param) {
		WriteErrorResponse(c, InvalidParam)
		return
	}

	// login
	if err := logic.Login(param); err != nil {
		WriteErrorResponse(c, UserNotExist)
	}

	WriteSuccessResponse(c, nil)
}
