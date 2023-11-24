package rdb

import (
	model "bluebell/Model"
	"math/rand"
	"testing"
)

func TestPost1(t *testing.T) {
	flushrdb()

	npost := 100
	posts := generatePost(npost)

	for _, post := range posts {
		if err := CreatePost(post); err != nil {
			t.Error(err.Error())
			return
		}
	}

	length, err := GetPostNumber()
	if err != nil {
		t.Error(err.Error())
	}
	if length != int64(npost) {
		t.Errorf("GetPostNumber fail, length: %d", length)
		return
	}

	pids, err := getPostIDs(npost)
	if err != nil {
		t.Error(err.Error())
	}

	for _, pid := range pids {
		i := 0
		for _, post := range posts {
			if post.ID == pid {
				continue
			}
			i++
		}
		if i == len(posts) {
			t.Errorf("pid %d not in posts", pid)
		}
	}

	return
}

// test getpostbycommunity
func TestPost2(t *testing.T) {
	flushrdb()

	npost := 100
	ncomm := 4
	var posts []*model.Post
	for i := 1; i <= ncomm; i++ {
		posts = append(posts, generatePostByCommunity(npost, i)...)
	}

	for _, post := range posts {
		if err := CreatePost(post); err != nil {
			t.Error(err.Error())
		}
	}

	length, err := GetPostNumber()
	if err != nil {
		t.Error(err.Error())
	}
	if length != int64(ncomm*npost) {
		t.Errorf("GetPostNumber fail, length: %d, need %d", length, ncomm*npost)
	}

	for cid := 1; cid <= ncomm; cid++ {
		pids, err := getPostIDsOfCommunity(npost, cid)
		if err != nil {
			t.Error(err.Error())
		}

		for _, pid := range pids {
			i := 0
			for _, post := range posts {
				if post.ID == pid && post.CommunityID == int64(cid) {
					continue
				}
				i++
			}
			if i == len(posts) {
				t.Errorf("pid %d not in posts", pid)
			}
		}
	}

	return
}

func generatePost(npost int) (posts []*model.Post) {
	posts = make([]*model.Post, 0)
	for i := 0; i < npost; i++ {
		posts = append(posts, &model.Post{
			ID:          int64(i),
			Title:       "title",
			Content:     "content",
			AuthorID:    int64(rand.Intn(100)),
			CommunityID: int64(1 + rand.Intn(4)),
		})
	}
	return
}

func generatePostByCommunity(npost int, cid int) (posts []*model.Post) {
	posts = make([]*model.Post, 0)
	for i := 0; i < npost; i++ {
		posts = append(posts, &model.Post{
			ID:          int64(cid*npost + i),
			Title:       "title",
			Content:     "content",
			AuthorID:    int64(rand.Intn(100)),
			CommunityID: int64(cid),
		})
	}
	return
}

func getPostIDs(npost int) (pids []int64, err error) {
	param := &model.ParamPostsQuary{
		Page:  1,
		Size:  int64(npost),
		Order: model.TIME,
	}

	pids, err = GetPostIDs(param)
	if err != nil {
		return
	}

	return
}

func getPostIDsOfCommunity(npost int, cid int) (pids []int64, err error) {
	param := &model.ParamPostsQuary{
		Page:        1,
		Size:        int64(npost),
		Order:       model.TIME,
		CommunityID: int64(cid),
	}

	pids, err = GetPostIDsOfCommunity(param)
	if err != nil {
		return
	}

	return
}

func flushrdb() {
	rdb.FlushAll()
}
