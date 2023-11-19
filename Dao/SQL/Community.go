package sql

import (
	log "bluebell/Log"
	model "bluebell/Model"
	"time"
)

var communities = []model.Community{
	// Go
	{
		ID:            1,
		Name:          "Go",
		Introducation: "Learn Golang with us",
		Create_time:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		Update_time:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	// AI
	{
		ID:            2,
		Name:          "AI",
		Introducation: "This is a community for AI, we are trying to make the world better",
		Create_time:   time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC),
		Update_time:   time.Date(2022, 6, 1, 0, 0, 0, 0, time.UTC),
	},
	// Honkai: Star Rail
	{
		ID:            3,
		Name:          "Honkai: Star Rail",
		Introducation: "Honkai: Star Rail is a game developed by miHoYo",
		Create_time:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		Update_time:   time.Date(2023, 4, 14, 0, 0, 0, 0, time.UTC),
	},
	// Sports
	{
		ID:            4,
		Name:          "Sports",
		Introducation: "Sports is good for health",
		Create_time:   time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		Update_time:   time.Date(2022, 8, 20, 13, 10, 0, 0, time.UTC),
	},
}

func InsertCommunity(community *model.Community) (err error) {
	log.Infof("Create community %v", community)
	err = db.Create(community).Error
	if err != nil {
		log.Errorf(err.Error())
	} else {
		log.Infof("Create community %v success", community)
	}
	return
}

func GetCommunities() (communities []model.Community, err error) {
	communities = make([]model.Community, 0)
	err = db.Find(&communities).Error
	for _, community := range communities {
		log.Infof("Get community %v", community)
	}
	if err != nil {
		log.Errorf(err.Error())
	} else {
		log.Infof("Get community list success")
	}

	return
}

func GetCommunity(id int64) (community model.Community, err error) {
	community = model.Community{}
	err = db.Where("id = ?", id).First(&community).Error
	if err != nil {
		log.Errorf(err.Error())
	} else {
		log.Infof("Get community %v success", community)
	}

	return
}
