package control

import (
	logic "bluebell/Logic"
	model "bluebell/Model"

	"github.com/gin-gonic/gin"
)

func VoteForPostHandler(ctx *gin.Context) {
	param := &model.ParamVote{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		WriteErrorResponse(ctx, InvalidParam)
		return
	}

	userid, err := GetUserID(ctx)
	if err != nil {
		WriteErrorResponse(ctx, InvalidToken)
		return
	}
	param.UserID = userid

	if err := logic.VoteForPost(param); err != nil {
		WriteErrorResponse(ctx, ServerBusy)
		return
	}

	WriteSuccessResponse(ctx, nil)
	return
}
