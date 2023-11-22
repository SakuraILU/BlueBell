package model

import "time"

type Community struct {
	ID            int64  `gorm:"primaryKey"`
	Name          string `gorm:"unique"`
	Introducation string
	Create_time   time.Time
	Update_time   time.Time
}

type ParamCommunity struct {
	ID   int64
	Name string
}

type ParamCommunityDetail struct {
	Name          string
	Introducation string
	Create_time   time.Time
	Update_time   time.Time
}
