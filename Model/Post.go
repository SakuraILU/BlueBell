package model

import "time"

type Post struct {
	ID          int64 `gorm:"primaryKey"`
	Title       string
	Content     string
	AuthorID    int64
	CommunityID int64
	Create_time time.Time
	Update_time time.Time
}

type ParamPost struct {
	Title       string
	Content     string
	AuthorID    int64
	CommunityID int64
}

type ParamPostDetail struct {
	AuthorName      string
	NVote           int64
	Post            *Post
	CommunityDetail *ParamCommunityDetail
}
