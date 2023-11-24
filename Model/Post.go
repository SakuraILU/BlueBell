package model

import "time"

type Post struct {
	// auto increment
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

const (
	TIME  = "time"
	SCORE = "score"
)

type ParamPostsQuary struct {
	CommunityID int64
	Page        int64
	Size        int64
	Order       string
}
