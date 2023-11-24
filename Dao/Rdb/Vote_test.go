package rdb

import (
	model "bluebell/Model"
	"math/rand"
	"testing"
)

func TestVote1(t *testing.T) {
	post_rdb.FlushAll()

	posts := generatePost(5)
	for _, post := range posts {
		if err := CreatePost(post); err != nil {
			t.Error(err.Error())
			return
		}
	}

	votes := generateVote(50, posts)
	for _, vote := range votes {
		if err := VoteForPost(vote); err != nil {
			t.Error(err.Error())
			return
		}
	}

	pid2vote := make(map[int64]int64)
	// cacluate positive vote for each post by myself
	for _, post := range posts {
		nvote := int64(0)
		for _, vote := range votes {
			t.Logf("vote: %v", vote)
			if vote.PostID == post.ID && vote.Choice == model.POSITIVE {
				nvote++
			}
		}
		pid2vote[post.ID] = nvote
	}

	// test get positive vote
	pid2vote2 := make(map[int64]int64)
	for _, post := range posts {
		nvote, err := GetPositiveVote(post.ID)
		if err != nil {
			t.Error(err.Error())
			return
		}
		pid2vote2[post.ID] = nvote
	}

	// compare with pid2vote
	for _, post := range posts {
		if pid2vote[post.ID] != pid2vote2[post.ID] {
			t.Errorf("pid %d positive vote not match, %d != %d", post.ID, pid2vote[post.ID], pid2vote2[post.ID])
		}
	}

	return
}

func TestVote2(t *testing.T) {
	post_rdb.FlushAll()

	posts := generatePost(5)
	for _, post := range posts {
		if err := CreatePost(post); err != nil {
			t.Error(err.Error())
			return
		}
	}

	votes := generateVote(50, posts)
	for _, vote := range votes {
		if err := VoteForPost(vote); err != nil {
			t.Error(err.Error())
			return
		}
	}

	pid2vote := make(map[int64]int64)
	// cacluate positive vote for each post by myself
	for _, post := range posts {
		nvote := int64(0)
		for _, vote := range votes {
			t.Logf("vote: %v", vote)
			if vote.PostID == post.ID && vote.Choice == model.POSITIVE {
				nvote++
			}
		}
		pid2vote[post.ID] = nvote
	}

	// test get positive votes
	pids := make([]int64, 0)
	for _, post := range posts {
		pids = append(pids, post.ID)
	}
	nvotes, err := GetPositiveVotes(pids)
	if err != nil {
		t.Error(err.Error())
		return
	}
	pid2vote2 := make(map[int64]int64)
	for i, pid := range pids {
		pid2vote2[pid] = nvotes[i]
	}

	// compare with pid2vote
	for _, post := range posts {
		if pid2vote[post.ID] != pid2vote2[post.ID] {
			t.Errorf("pid %d positive vote not match, %d != %d", post.ID, pid2vote[post.ID], pid2vote2[post.ID])
		}
	}

	return
}

func generateVote(nvote int, posts []*model.Post) []*model.ParamVote {
	votes := make([]*model.ParamVote, 0)
	for i := 0; i < nvote; i++ {
		vote := &model.ParamVote{
			UserID: int64(i),
			PostID: posts[rand.Intn(len(posts))].ID,
			Choice: int64(-1 + rand.Intn(3)),
		}
		votes = append(votes, vote)
	}

	return votes
}
