package control

import (
	rdb "bluebell/Dao/Rdb"
	cookie "bluebell/Utils/Cookie"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	KeyUserID   = "UserID"
	KeyUsername = "Username"
)

func JwtAuthorization() func(*gin.Context) {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")
		if auth == "" {
			ctx.Next()
			return
		}

		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			WriteErrorResponse(ctx, InvalidParam)
			ctx.Abort()
			return
		}

		token_str := parts[1]
		user, err := cookie.ParseToken(token_str)
		if err != nil {
			WriteErrorResponse(ctx, InvalidToken)
			ctx.Abort()
			return
		}

		// limit the number that a user can logged in at the same time
		if !rdb.TokenExist(user.ID, token_str) {
			WriteErrorResponse(ctx, InvalidToken)
			ctx.Abort()
			return
		}

		ctx.Set(KeyUserID, user.ID)
		ctx.Set(KeyUsername, user.Username)

		ctx.Next()
		return
	}
}

func GetUserID(ctx *gin.Context) (id int64, err error) {
	v, ok := ctx.Get(KeyUserID)
	if !ok {
		err = fmt.Errorf("User not login yet")
		return
	}
	id = v.(int64)

	return
}
