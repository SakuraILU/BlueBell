package sql

import (
	log "bluebell/Log"
	model "bluebell/Model"
)

func CreatePost(post *model.Post) (err error) {
	err = db.Create(post).Error
	if err != nil {
		log.Errorf(err.Error())
	} else {
		log.Infof("Insert post %v success", post)
	}

	return
}

func GetPosts(page, size int) (posts []*model.Post, err error) {
	posts = make([]*model.Post, 0)
	err = db.Offset((page - 1) * size).Limit(size).Find(&posts).Error
	if err != nil {
		log.Errorf(err.Error())
	} else {
		log.Infof("Get %d posts success", len(posts))
	}

	return
}

func GetPostByID(id int64) (post *model.Post, err error) {
	post = &model.Post{}
	err = db.Where("id = ?", id).First(post).Error
	if err != nil {
		log.Errorf(err.Error())
	} else {
		log.Infof("Get post %v success", post.ID)
	}

	return
}