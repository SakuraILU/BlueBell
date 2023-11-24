package rdb

import (
	log "bluebell/Log"
	model "bluebell/Model"
	"fmt"
	"strconv"
	"time"
)

const (
	TTL_POST_INORDER_OF_COMMUNITY = 20
)

func CreatePost(post *model.Post) (err error) {
	post_of_community_key := fmt.Sprintf("%s%d", KEYPOST_OF_COMMUNITY_ZSET, post.CommunityID)
	rdb.EvalSha(scripts[SETPOST].Sha, []string{KEYPOST_SCORE_ZSET, KEYPOST_TIME_ZSET, post_of_community_key}, post.ID, time.Now().Unix())
	if err != nil {
		log.Errorf("Create post %v fail", post)
	}

	return err
}

func GetPostIDs(param *model.ParamPostsQuary) (pids []int64, err error) {
	key := KEYPOST_TIME_ZSET
	if param.Order == model.SCORE {
		key = KEYPOST_SCORE_ZSET
	}

	start := (param.Page - 1) * param.Size
	stop := param.Page*param.Size - 1
	res, err := rdb.ZRevRange(key, start, stop).Result()
	if err != nil {
		return
	}

	pids = make([]int64, 0)
	for _, v := range res {
		pid, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		pids = append(pids, int64(pid))
	}

	return
}

func GetPostIDsOfCommunity(param *model.ParamPostsQuary) (pids []int64, err error) {
	key_post_of_community := fmt.Sprintf("%s%d", KEYPOST_OF_COMMUNITY_ZSET, param.CommunityID)

	key_post_inorder := KEYPOST_TIME_ZSET
	key_post_inorder_of_community := fmt.Sprintf("%s%d", KEYPOST_TIME_OF_COMMUNITY, param.CommunityID)
	if param.Order == model.SCORE {
		key_post_inorder = KEYPOST_SCORE_ZSET
		key_post_inorder_of_community = fmt.Sprintf("%s%d", KEYPOST_SCORE_OF_COMMUNITY, param.CommunityID)
	}

	start := (param.Page - 1) * param.Size
	stop := param.Page*param.Size - 1
	// print redis eval keys and args
	res, err := rdb.EvalSha(scripts[GETPOSTOFCOMMUNITY].Sha, []string{key_post_inorder, key_post_of_community, key_post_inorder_of_community}, TTL_POST_INORDER_OF_COMMUNITY, start, stop).Result()

	pids = make([]int64, 0)
	for _, v := range res.([]interface{}) {
		pid, err := strconv.Atoi(v.(string))
		if err != nil {
			return nil, err
		}
		pids = append(pids, int64(pid))
	}

	return
}

func GetPostNumber() (n int64, err error) {
	n, err = rdb.ZCard(KEYPOST_SCORE_ZSET).Result()
	if err != nil {
		log.Errorf("Get post number fail")
	}

	return
}
