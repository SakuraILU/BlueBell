package rdb

import (
	config "bluebell/Config"

	"github.com/go-redis/redis"
)

const (
	KEYTOKEN_USER_OF_SET       = "token_of_user"             // token of user_uid
	KEYPOST_SCORE_ZSET         = "(post:score)"              // post:score
	KEYPOST_TIME_ZSET          = "(post:time)"               // post:time
	KEYUSER_VOTE_OF_POST_ZSET  = "(user:vote)_of_post"       // (user:vote) of post_pid
	KEYPOST_OF_COMMUNITY_ZSET  = "post_of_community"         // posts of community_cid
	KEYPOST_TIME_OF_COMMUNITY  = "(post:time)_of_community"  // post:time of community_cid
	KEYPOST_SCORE_OF_COMMUNITY = "(post:score)_of_community" // post:score of community_cid
)

var (
	token_rdb *redis.Client
	// post_rdb  *redis.Client
	post_rdb *redis.Client
)

func init() {
	token_rdb = redis.NewClient(&redis.Options{
		Addr:     config.Cfg.Redis.Addr,
		Password: config.Cfg.Redis.Password,
		DB:       config.Cfg.Redis.TokenDB,
	})

	post_rdb = redis.NewClient(&redis.Options{
		Addr:     config.Cfg.Redis.Addr,
		Password: config.Cfg.Redis.Password,
		DB:       config.Cfg.Redis.PostDB,
	})

	// all the database shares their scripts
	// so just use any one of them to load scripts
	for i, script := range scripts {
		sha, err := token_rdb.ScriptLoad(script.Lua).Result()
		if err != nil {
			panic(err)
		}
		scripts[i].Sha = sha
	}
}
