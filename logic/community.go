package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

func GetCommunityList() (communityList []*models.CommunityCategoryList, err error) {
	// 从数据库中获取
	return mysql.GetCommunityCategoryList()
}

func GetCommunityDetailByID(communityID int64) (*models.CommunityCategoryDetail, error) {
	return mysql.GetCommunityDetailByID(communityID)
}
