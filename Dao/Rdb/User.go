package rdb

import (
	config "bluebell/Config"
	log "bluebell/Log"
	"fmt"
)

func SetToken(userid int64, token_str string) (err error) {
	nduplogin := config.Cfg.Logic.NDuplicateLogin
	key := fmt.Sprintf("%s%d", KEYTOKEN_USER_OF_SET, userid)
	if _, err = token_rdb.EvalSha(scripts[SETTOKEN].Sha, []string{key}, token_str, nduplogin).Result(); err != nil {
		log.Errorf(err.Error())
		return
	}

	log.Infof("Set user %v's token %v success", userid, token_str)

	return
}

func GetTokenStrs(userid int64) (token_strs []string, err error) {
	key := fmt.Sprintf("%s%d", KEYTOKEN_USER_OF_SET, userid)
	if token_strs, err = token_rdb.LRange(key, 0, -1).Result(); err != nil {
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
