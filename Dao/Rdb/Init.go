package rdb

import (
	log "bluebell/Log"

	"github.com/go-redis/redis"
)

const (
	KEYTOKEN   = "token_user"
	NDUPLICATE = 3
)

var (
	rdb                 *redis.Client
	settoken_script_sha string
)

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	settoken_script := `
		local key = KEYS[1]
		local token_str = ARGV[1]
		local cap = tonumber(ARGV[2])

		while redis.call("LLEN", key) >= cap do
			redis.call("LPOP", key)
		end

		redis.call("RPUSH", key, token_str)

		return 1
	`

	var err error
	settoken_script_sha, err = rdb.ScriptLoad(settoken_script).Result()
	if err != nil {
		log.Panic(err.Error())
	}

	log.Infof("Setting Redis OK")
}
