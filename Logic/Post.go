package logic

import (
	rdb "bluebell/Dao/Rdb"
	sql "bluebell/Dao/SQL"
	log "bluebell/Log"
	model "bluebell/Model"
	"fmt"
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

func GetPosts(param *model.ParamPostsQuary) (p_details []*model.ParamPostDetail, err error) {
	npost, err := rdb.GetPostNumber()
	if err != nil {
		return nil, err
	}
	start := (param.Page - 1) * param.Size
	if start >= npost {
		err = fmt.Errorf("page %d is out of range", param.Page)
		log.Errorf(err.Error())
		return
	}

	var pids []int64
	if param.CommunityID == 0 {
		pids, err = rdb.GetPostIDs(param)
	} else {
		if _, err := sql.GetCommunityByID(param.CommunityID); err != nil {
			log.Errorf(err.Error())
			return nil, err
		}
		pids, err = rdb.GetPostIDsOfCommunity(param)
	}

	if err != nil {
		return
	}

	p_details = make([]*model.ParamPostDetail, 0)

	posts, err := sql.GetPostsByIDs(pids)
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
