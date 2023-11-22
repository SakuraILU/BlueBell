package logic

import (
	rdb "bluebell/Dao/Rdb"
	sql "bluebell/Dao/SQL"
	model "bluebell/Model"
	"fmt"
)

func VoteForPost(vote *model.ParamVote) (err error) {
	if _, err = sql.GetPostByID(vote.PostID); err != nil {
		return fmt.Errorf("post id %d not exist", vote.PostID)
	}

	err = rdb.VoteForPost(vote)
	return
}
