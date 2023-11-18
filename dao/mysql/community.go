package mysql

import "bluebell/models"

func GetCommunityCategoryList() (communityList []*models.CommunityCategoryList, err error) {
	sqlStr := "select community_id,community_name from community"
	if err = db.Select(&communityList, sqlStr); err != nil {
		return communityList, err
	}
	return
}

func GetCommunityDetailByID(communityId int64) (communityCategoryDetail *models.CommunityCategoryDetail, err error) {
	sqlStr := "select community_id,community_name,introduction from community where community_id = ?"
	// 将空指针进行初始化，指向一片区域
	communityCategoryDetail = new(models.CommunityCategoryDetail)
	err = db.Get(communityCategoryDetail, sqlStr, communityId)
	return
}
