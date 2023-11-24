package sql

import (
	log "bluebell/Log"
	model "bluebell/Model"
)

func CreateUser(user *model.User) (err error) {
	_, err = GetUserByName(user.Username)
	if err == nil {
		log.Errorf("User %v already exist", user)
		return
	}

	err = db.Create(&user).Error
	if err != nil {
		log.Errorf(err.Error())
	} else {
		log.Infof("Regist user %v success", user)
	}

	return
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
