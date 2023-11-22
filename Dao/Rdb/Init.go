package rdb

import (
	"github.com/go-redis/redis"
)

const (
	NDUPLICATE                    = 3
	KEYTOKEN_USER_SET_PREFIX      = "token_user"
	KEYPOST_SCORE_ZSET            = "(post:score)"
	KEYPOST_TIME_ZSET             = "(post:time)"
	KEYUSER_VOTE_POST_ZSET_PREFIX = "(user:vote)_post"
)

var (
	rdb *redis.Client
)

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	for _, script := range scripts {
		sha, err := rdb.ScriptLoad(script.Lua).Result()
		if err != nil {
			panic(err)
		}
		script.Sha = sha
	}
}
