package control

import (
	log "bluebell/Log"
	logic "bluebell/Logic"
	model "bluebell/Model"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(ctx *gin.Context) {
	param := &model.ParamSignUp{}

	if err := ctx.ShouldBindJSON(param); err != nil {
		log.Error(err)
		WriteErrorResponse(ctx, InvalidParam)
		return
	}

	if !validateSignUp(param) {
		WriteErrorResponse(ctx, InvalidParam)
		return
	}

	log.Warnf("try to regist user %v", param)

	// signup the user
	if err := logic.SignUp(param); err != nil {
		WriteErrorResponse(ctx, UserExist)
	} else {
		WriteSuccessResponse(ctx, nil)
	}
}

func LoginHandler(ctx *gin.Context) {
	// check fast login through cookie...
	if _, err := GetUserID(ctx); err == nil {
		WriteSuccessResponse(ctx, nil)
		return
	}

	param := &model.ParamLogin{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		log.Error(err)
		WriteErrorResponse(ctx, InvalidParam)
		return
	}

	if !validateLogin(param) {
		WriteErrorResponse(ctx, InvalidParam)
		return
	}

	// login
	if token_str, err := logic.Login(param); err == nil {
		WriteSuccessResponse(ctx, token_str)
	} else {
		WriteErrorResponse(ctx, UserNotExist)
	}

	return
}
