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

	return err
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

	posts = sortOrderByPID(posts, pids)

	// for i, pid := range pids {
	// 	if posts[i].ID != pid {
	// 		log.Panic("post.ID is not match pid...")
	// 	}
	// }
	return
}

func sortOrderByPID(posts []*model.Post, pids []int64) []*model.Post {
	for i, pid := range pids {
		if posts[i].ID == pid {
			continue
		}

		j := i + 1
		for ; j < len(pids); j++ {
			if posts[j].ID == pid {
				break
			}
		}

		tmp := posts[i]
		posts[i] = posts[j]
		posts[j] = tmp
	}

	return posts
}
