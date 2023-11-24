package control

import (
	config "bluebell/Config"
	ratelimit "bluebell/Utils/RateLimit"
	"time"

	"github.com/gin-gonic/gin"
)

func RateLimit() func(*gin.Context) {
	rate, nbucket := config.Cfg.RateLimit.Rate, config.Cfg.RateLimit.NBucket
	limiter := ratelimit.NewRateLimit(rate, nbucket)

	return func(ctx *gin.Context) {
		for !limiter.Allow() {
			time.Sleep(300 * time.Millisecond)
		}

		ctx.Next()
		return
	}
}
