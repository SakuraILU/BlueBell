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
	cid_str, page_str, size_str, order := ctx.Query("cid"), ctx.Query("page"), ctx.Query("size"), ctx.Query("order")
	cid, err := strconv.Atoi(cid_str)
	if err != nil {
		WriteErrorResponse(ctx, InvalidParam)
		return
	}
	page, err := strconv.Atoi(page_str)
	if err != nil || page < 1 {
		WriteErrorResponse(ctx, InvalidParam)
		return
	}
	size, err := strconv.Atoi(size_str)
	if err != nil || size < 1 {
		WriteErrorResponse(ctx, InvalidParam)
		return
	}

	param := &model.ParamPostsQuary{
		CommunityID: int64(cid),
		Page:        int64(page),
		Size:        int64(size),
		Order:       order,
	}

	if !validatePostsQuary(param) {
		WriteErrorResponse(ctx, InvalidParam)
		return
	}

	// logic.GetPosts
	posts, err := logic.GetPosts(param)
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
}
