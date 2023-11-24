package ratelimit

import (
	config "bluebell/Config"
	log "bluebell/Log"
	"time"

	"github.com/go-redis/redis"
)

const (
	KEYNTOKEN   = "ntoken"
	KEYLASTTIME = "lasttime"
)

type RateLimit struct {
	rdb        *redis.Client
	sha_script string
	rate       int
	nbucket    int
}

func NewRateLimit(rate int, nbucket int) *RateLimit {
	rl := &RateLimit{
		rdb: redis.NewClient(&redis.Options{
			Addr:     config.Cfg.Redis.Addr,
			Password: config.Cfg.Redis.Password,
			DB:       config.Cfg.Redis.RateLimitDB,
		}),
		rate:    rate,
		nbucket: nbucket,
	}

	sha, err := rl.rdb.ScriptLoad(script).Result()
	if err != nil {
		log.Panic(err.Error())
	}
	rl.sha_script = sha

	rl.rdb.Set(KEYNTOKEN, 0, 0)
	// second time
	curtime := time.Now().Unix()
	rl.rdb.Set(KEYLASTTIME, curtime, 0)

	return rl
}

func (rl *RateLimit) AllowN(n int) bool {
	curtime := time.Now().Unix()
	res, err := rl.rdb.EvalSha(rl.sha_script, []string{KEYNTOKEN, KEYLASTTIME}, n, rl.rate, rl.nbucket, curtime).Result()

	if err != nil {
		return false
	}

	return res.(int64) == 1
}

func (rl *RateLimit) Allow() bool {
	return rl.AllowN(1)
}
