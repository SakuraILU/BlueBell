package sql

import (
	log "bluebell/Log"
	model "bluebell/Model"

	"gorm.io/gorm"
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

func GetPostsByIDs(pids []int64) (posts []*model.Post, err error) {
	posts = make([]*model.Post, 0)
	err = db.Where("id in (?)", pids).Order(gorm.Expr("FIELD(id, ?)", pids)).Find(&posts).Error
	if err != nil {
		log.Errorf(err.Error())
	} else {
		log.Infof("Get %d posts success", len(posts))
	}

	return
}
