package model

const (
	NEGTIVE  = "-1"
	NOVATE   = "0"
	POSITIVE = "1"
)

type ParamVote struct {
	UserID int64
	PostID int64
	Choice int64 // 1 for agree, 0 for nothing, -1 for disagree
}
