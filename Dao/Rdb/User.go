package rdb

import (
	log "bluebell/Log"
	"fmt"

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

func SetToken(userid int64, token_str string) (err error) {
	key := fmt.Sprintf("%s%d", KEYTOKEN, userid)
	if _, err = rdb.EvalSha(settoken_script_sha, []string{key}, token_str, NDUPLICATE).Result(); err != nil {
		log.Errorf(err.Error())
		return
	}

	log.Infof("Set user %v's token %v success", userid, token_str)

	return
}

func GetTokenStrs(userid int64) (token_strs []string, err error) {
	key := fmt.Sprintf("%s%d", KEYTOKEN, userid)
	if token_strs, err = rdb.LRange(key, 0, -1).Result(); err != nil {
		log.Errorf("Get user %v's tokens fail", userid)
		return
	}

	return
}

func TokenExist(userid int64, token_str string) (exist bool) {
	token_strs, err := GetTokenStrs(userid)
	if err != nil {
		exist = false
		log.Errorf("Get user %v's token fail", userid)
		return
	}

	for _, v := range token_strs {
		if v == token_str {
			exist = true
			return
		}
	}

	return
}
