package logic

import (
	config "bluebell/Config"
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
	starttime := config.Cfg.Snowflake.StartTime
	datacenterid, machineid := config.Cfg.Snowflake.DataCenterID, config.Cfg.Snowflake.MachineID
	stime, err := time.Parse("2006-01-02 15:04:05", starttime)
	if err != nil {
		log.Panic(err.Error())
	}
	user_sf, err = snowflake.NewSnowflake(stime, datacenterid, machineid)
	if err != nil {
		log.Panic(err.Error())
	}
	post_sf, err = snowflake.NewSnowflake(stime, datacenterid, machineid)
	if err != nil {
		log.Panic(err.Error())
	}
}
