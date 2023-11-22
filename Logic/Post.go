package logic

import (
	rdb "bluebell/Dao/Rdb"
	sql "bluebell/Dao/SQL"
	model "bluebell/Model"
	"time"
)

func CreatePost(param *model.ParamPost) (err error) {
	post := &model.Post{
		ID:          post_sf.NextID(),
		Title:       param.Title,
		Content:     param.Content,
		AuthorID:    param.AuthorID,
		CommunityID: param.CommunityID,
		Create_time: time.Now(),
		Update_time: time.Now(),
	}

	_, err = sql.GetCommunityByID(param.CommunityID)
	if err != nil {
		return
	}

	if err = sql.CreatePost(post); err != nil {
		return
	}

	if err = rdb.CreatePost(post); err != nil {
		return
	}

	return
}

func GetPosts(page, size int) (p_details []*model.ParamPostDetail, err error) {
	p_details = make([]*model.ParamPostDetail, 0)

	posts, err := sql.GetPosts(page, size)
	if err != nil {
		return nil, err
	}

	for _, post := range posts {
		var user *model.User
		user, err = sql.GetUserByID(post.AuthorID)
		if err != nil {
			return
		}

		var community *model.ParamCommunityDetail
		community, err = GetCommunityDetail(post.CommunityID)
		if err != nil {
			return
		}

		if len(post.Content) > posts_len {
			post.Content = post.Content[0:posts_len] + ignore_sign
		}

		p_detail := &model.ParamPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}

		p_details = append(p_details, p_detail)
	}

	pids := make([]int64, 0)
	for _, post := range posts {
		pids = append(pids, post.ID)
	}
	nvotes, err := rdb.GetPositiveVotes(pids)
	if err != nil {
		return nil, err
	}
	for i, p_detail := range p_details {
		p_detail.NVote = nvotes[i]
	}

	return
}

func GetPost(id int64) (param_post *model.ParamPostDetail, err error) {
	post, err := sql.GetPostByID(id)
	if err != nil {
		return
	}

	user, err := sql.GetUserByID(post.AuthorID)
	if err != nil {
		return
	}

	community, err := sql.GetCommunityByID(post.CommunityID)
	if err != nil {
		return
	}
	c_detail := &model.ParamCommunityDetail{
		Name:          community.Name,
		Introducation: community.Introducation,
		Create_time:   community.Create_time,
		Update_time:   community.Update_time,
	}

	nvote, err := rdb.GetPositiveVote(post.ID)
	if err != nil {
		return
	}

	param_post = &model.ParamPostDetail{
		AuthorName:      user.Username,
		NVote:           nvote,
		Post:            post,
		CommunityDetail: c_detail,
	}

	return
}
