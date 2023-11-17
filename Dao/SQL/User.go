package sql

import (
	log "bluebell/Log"
	model "bluebell/Model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("bluebell.db"), &gorm.Config{})
	if err != nil {
		log.Panic(err.Error())
	}
	db.AutoMigrate(&model.User{})
}

func InsertUser(user *model.User) (err error) {
	log.Errorf("Regist user %v", user)
	err = db.Create(user).Error
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
	user = new(model.User)
	err = db.Where("username = ?", name).First(user).Error
	if err != nil {
		log.Errorf(err.Error())
	} else {
		log.Infof("Get user %v success", user)
	}

	return
}
