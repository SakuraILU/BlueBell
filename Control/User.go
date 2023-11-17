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

func validateSignUp(param *model.ParamSignUp) bool {
	if len(param.Username) <= 0 || len(param.Password) < 6 || len(param.RePassword) < 6 {
		return false
	}

	if param.Password != param.RePassword {
		return false
	}

	return true
}
