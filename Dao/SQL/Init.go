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

	db.AutoMigrate(&model.Community{})
	// check if the communities is empty
	var count int64
	db.Model(&model.Community{}).Count(&count)
	if count == 0 {
		for _, community := range communities {
			if err := CreateCommunity(&community); err != nil {
				panic(err)
			}
		}
	}

	db.AutoMigrate(&model.Post{})
}
