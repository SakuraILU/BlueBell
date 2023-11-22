package control

import (
	logic "bluebell/Logic"
	model "bluebell/Model"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	ParamPID = "pid"
)

func CreatePostHandler(ctx *gin.Context) {
	param := &model.ParamPost{}
	if err := ctx.ShouldBindJSON(param); err != nil {
		WriteErrorResponse(ctx, InvalidParam)
		return
	}

	userid, err := GetUserID(ctx)
	if err != nil {
		WriteErrorResponse(ctx, InvalidToken)
		return
	}
	param.AuthorID = userid

	// logic.CreatePost
	err = logic.CreatePost(param)
	if err != nil {
		WriteErrorResponse(ctx, ServerBusy)
	}

	WriteSuccessResponse(ctx, nil)
}

func GetPostListHandler(ctx *gin.Context) {
	page_str, size_str := ctx.Query("page"), ctx.Query("size")
	page, err := strconv.Atoi(page_str)
	if err != nil {
		WriteErrorResponse(ctx, InvalidParam)
		return
	}
	size, err := strconv.Atoi(size_str)
	if err != nil {
		WriteErrorResponse(ctx, InvalidParam)
		return
	}

	// logic.GetPosts
	posts, err := logic.GetPosts(page, size)
	if err != nil {
		WriteErrorResponse(ctx, ServerBusy)
		return
	}

	WriteSuccessResponse(ctx, posts)
}

func GetPostDetailHandler(ctx *gin.Context) {
	id_str := ctx.Param(ParamPID)
	id, err := strconv.Atoi(id_str)
	if err != nil {
		WriteErrorResponse(ctx, InvalidParam)
		return
	}

	// logic.GetPost
	post, err := logic.GetPost(int64(id))
	if err != nil {
		WriteErrorResponse(ctx, PostNotExist)
		return
	}

	WriteSuccessResponse(ctx, post)
	return
}
