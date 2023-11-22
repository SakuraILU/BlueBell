package sql

import (
	log "bluebell/Log"
	model "bluebell/Model"
)

func CreateUser(user *model.User) (err error) {
	err = db.Create(&user).Error
	if err != nil {
		log.Errorf(err.Error())
	} else {
		log.Infof("Regist user %v success", user)
	}
	return
}

func CheckUserExistByName(name string) bool {
	if _, err := GetUserByName(name); err == nil {
		return true
	} else {
		return false
	}
}

func GetUserByName(name string) (user *model.User, err error) {
	user = &model.User{}
	err = db.Where("username = ?", name).First(user).Error
	if err != nil {
		log.Errorf(err.Error())
	} else {
		log.Infof("Get user %v success", user)
	}

	return
}

func GetUserByID(id int64) (user *model.User, err error) {
	user = &model.User{}
	err = db.Where("id = ?", id).First(user).Error
	if err != nil {
		log.Errorf(err.Error())
	} else {
		log.Infof("Get user %v success", user)
	}

	return
}
