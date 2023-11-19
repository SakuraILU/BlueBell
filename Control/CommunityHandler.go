package control

import (
	log "bluebell/Log"
	logic "bluebell/Logic"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CommunityListHandler(ctx *gin.Context) {
	// logic.GetCommunityList
	param_communities, err := logic.GetCommunities()
	if err != nil {
		WriteErrorResponse(ctx, ServerBusy)
		return
	}

	WriteSuccessResponse(ctx, param_communities)
	return
}

func CommunityDetail(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Errorf("Invalid community id")
		WriteErrorResponse(ctx, InvalidParam)
		return
	}

	// logic.GetCommunityDetail
	c_detail, err := logic.GetCommunityDetail(int64(id))
	if err != nil {
		WriteErrorResponse(ctx, CommunityNotExist)
		return
	}

	WriteSuccessResponse(ctx, c_detail)
	return
}
