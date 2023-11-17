package logic

import (
	sql "bluebell/Dao/SQL"
	log "bluebell/Log"
	model "bluebell/Model"
	snowflake "bluebell/Utils/Snowflake"
	"crypto/md5"
	"fmt"
	"time"
)

var (
	user_sf *snowflake.Snowflake
	salt    string = "honkai: star rail"
)

func init() {
	stime, err := time.Parse("2006-01-02 15:04:05", "2020-01-01 00:00:00")
	if err != nil {
		log.Panic(err.Error())
	}
	user_sf, err = snowflake.NewSnowflake(stime, 0, 0)
	if err != nil {
		log.Panic(err.Error())
	}
}

func SignUp(param *model.ParamSignUp) (err error) {
	user := model.User{
		ID:       user_sf.NextID(),
		Username: param.Username,
		Password: encryptPassword(param.Password),
	}

	if sql.CheckUserExistByName(user.Username) {
		err = fmt.Errorf("user %s already exist", user.Username)
		return
	}
	// dao: write to database
	if err = sql.InsertUser(&user); err != nil {
		return
	}

	return
}

func encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(password))
	epassword := h.Sum([]byte(salt))

	return string(epassword)
}

func Login(param *model.ParamLogin) (err error) {
	var user *model.User

	if user, err = sql.GetUserByName(param.Username); err != nil {
		err = fmt.Errorf("user %s not exist", user.Username)
		return
	}

	if user.Password != encryptPassword(param.Password) {
		err = fmt.Errorf("password is not correct")
		return
	}

	log.Infof("user %s login success", user.Username)

	return
}
