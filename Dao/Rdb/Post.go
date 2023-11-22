package rdb

import (
	log "bluebell/Log"
	model "bluebell/Model"
	"time"
)

func CreatePost(post *model.Post) (err error) {
	rdb.EvalSha(scripts[SETPOST].Sha, []string{KEYPOST_SCORE_ZSET, KEYPOST_TIME_ZSET}, post.ID, time.Now().Unix())
	if err != nil {
		log.Errorf("Create post %v fail", post)
	}

	return err
}
