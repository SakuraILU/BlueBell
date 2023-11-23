package rdb

import (
	log "bluebell/Log"
	model "bluebell/Model"
	"fmt"
)

func VoteForPost(vote *model.ParamVote) (err error) {
	user_vote_of_post_key := fmt.Sprintf("%s%d", KEYUSER_VOTE_OF_POST_ZSET, vote.PostID)
	uid := fmt.Sprint(vote.UserID)
	orgchoice_float := rdb.ZScore(user_vote_of_post_key, uid).Val()
	orgchoice := int64(orgchoice_float)
	log.Errorf("from choice %v to choice %v", orgchoice, vote.Choice)
	if orgchoice == vote.Choice {
		return
	}

	if err = rdb.EvalSha(scripts[SETVOTE].Sha, []string{KEYPOST_SCORE_ZSET, user_vote_of_post_key}, vote.PostID, vote.UserID, orgchoice, vote.Choice).Err(); err != nil {
		return
	}

	return
}

func GetPositiveVote(pid int64) (nvote int64, err error) {
	key := fmt.Sprintf("%s%d", KEYUSER_VOTE_OF_POST_ZSET, pid)
	nvote, err = rdb.ZCount(key, model.POSITIVE, model.POSITIVE).Result()
	if err != nil {
		log.Errorf("Get positive vote for post %d fail", pid)
	}

	return
}

func GetPositiveVotes(pids []int64) (nvotes []int64, err error) {
	nvotes = make([]int64, 0)

	keys := make([]string, 0)
	for _, pid := range pids {
		keys = append(keys, fmt.Sprintf("%s%d", KEYUSER_VOTE_OF_POST_ZSET, pid))
	}
	log.Warn(keys)

	val, err := rdb.EvalSha(scripts[GETVOTES].Sha, keys, model.POSITIVE).Result()
	if err != nil {
		log.Errorf("Get positive votes for posts %v fail", pids)
		return
	}

	for _, nvote := range val.([]interface{}) {
		nvotes = append(nvotes, nvote.(int64))
	}

	return
}
