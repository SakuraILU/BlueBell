package logic

import (
	log "bluebell/Log"
	snowflake "bluebell/Utils/Snowflake"
	"time"
)

const (
	salt        string = "honkai: star rail"
	posts_len   int    = 128
	ignore_sign string = "..."
)

var (
	user_sf *snowflake.Snowflake
	post_sf *snowflake.Snowflake
)

func init() {
	stime, err := time.Parse("2006-01-02 15:04:05", "2020-01-01 00:00:00")
	if err != nil {
		log.Panic(err.Error())
	}
	user_sf, err = snowflake.NewSnowflake(stime, 0, 0)
	if err != nil {
		log.Panic(err.Error())
	}
	post_sf, err = snowflake.NewSnowflake(stime, 0, 0)
	if err != nil {
		log.Panic(err.Error())
	}
}
