package logic

import (
	sql "bluebell/Dao/SQL"
	model "bluebell/Model"
)

func GetCommunities() (param_communities []*model.ParamCommunity, err error) {
	param_communities = make([]*model.ParamCommunity, 0)
	communities, err := sql.GetCommunities()
	if err != nil {
		return
	}

	for _, community := range communities {
		param_communities = append(param_communities, &model.ParamCommunity{
			ID:   community.ID,
			Name: community.Name,
		})
	}

	return
}

func GetCommunityDetail(id int64) (c_detail *model.ParamCommunityDetail, err error) {
	community, err := sql.GetCommunityByID(id)
	if err != nil {
		return
	}

	c_detail = &model.ParamCommunityDetail{
		Name:          community.Name,
		Introducation: community.Introducation,
		Create_time:   community.Create_time,
		Update_time:   community.Update_time,
	}

	return
}
