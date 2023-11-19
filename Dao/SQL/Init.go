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

	// for _, community := range communities {
	// 	if err := InsertCommunity(&community); err != nil {
	// 		panic(err)
	// 	}
	// }
}
