package logic

import (
	rdb "bluebell/Dao/Rdb"
	sql "bluebell/Dao/SQL"
	log "bluebell/Log"
	model "bluebell/Model"
	cookie "bluebell/Utils/Cookie"
	"crypto/md5"
	"fmt"
)

func SignUp(param *model.ParamSignUp) (err error) {
	user := &model.User{
		ID:       user_sf.NextID(),
		Username: param.Username,
		Password: encryptPassword(param.Password),
	}

	// dao: write to database
	if err = sql.CreateUser(user); err != nil {
		return
	}

	return
}

func encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(password))
	h.Write([]byte(salt))

	return string(h.Sum(nil))
}

func Login(param *model.ParamLogin) (token_str string, err error) {
	user, err := sql.GetUserByName(param.Username)
	if err != nil {
		err = fmt.Errorf("user %s not exist", user.Username)
		return
	}

	if user.Password != encryptPassword(param.Password) {
		err = fmt.Errorf("password is not correct")
		return
	}

	token_str, err = cookie.GetToken(user)
	if err != nil {
		log.Errorf("cookie generation wrong!")
	}

	rdb.SetToken(user.ID, token_str)
	return token_str, nil
}
