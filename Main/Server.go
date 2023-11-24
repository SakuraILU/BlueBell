package main

import (
	config "bluebell/Config"
	control "bluebell/Control"

	"github.com/gin-gonic/gin"
)

var serveraddr = config.Cfg.IP + ":" + config.Cfg.Port

func startServer() {
	r := gin.Default()
	v1 := r.Group("/api/v1")

	v1.Use(control.RateLimit(), control.JwtAuthorization())

	v1.POST("/signup", control.SignUpHandler)
	v1.POST("/login", control.LoginHandler)
	v1.GET("/community", control.CommunityListHandler)
	v1.GET("/community/:cid", control.CommunityDetailHandler)
	v1.POST("/post", control.CreatePostHandler)
	v1.GET("/posts", control.GetPostListHandler)
	v1.GET("/post/:pid", control.GetPostDetailHandler)
	v1.POST("/vote", control.VoteForPostHandler)

	r.SetTrustedProxies([]string{"127.0.0.1"})

	r.Run(serveraddr)
}
