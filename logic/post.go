package logic

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/models"
)

func CreatePost(postContent *models.PostContent) error {
	// 将帖子的创建时间插入redis数据库，
	if err := redis.CreatePostTime(postContent.ID); err != nil {
		return err
	}
	// 插入mysql数据库
	return mysql.InsetNewPostContent(postContent)
}

func GetPostDetailByID(postID int64) (apiPostDetail *models.ApiPostDetail, err error) {
	apiPostDetail = new(models.ApiPostDetail)
	apiPostDetail.PostContent, err = mysql.GetPostDetailByID(postID)
	apiPostDetail.CommunityCategoryDetail, err = mysql.GetCommunityDetailByID(apiPostDetail.PostContent.CommunityID)
	apiPostDetail.AuthorName, err = mysql.GetAuthorNameByUserID(apiPostDetail.PostContent.AuthorID)
	return
}

func GetPostList(page int, size int) ([]*models.PostContent, error) {
	// mysql查寻
	return mysql.GetPostList(page, size)
}

func GetPostIDListOrderByTime(p *models.ParamPost) (apiPostDetailList []*models.ApiPostDetail, err error) {
	// 从redis中拿到按时间或分数排序的顺序的post_id
	postIDList, err := redis.GetPostIDListOrderByTimeOrScore(p)
	if err != nil {
		return nil, err
	}
	// 根据post_id查询每个帖子的投票数
	postVoteNumList, err := redis.GetPostVoteNumByPostID(postIDList)
	if err != nil {
		return nil, err
	}
	// 根据postIDList从Mysql中查询postList
	postList, err := mysql.GetPostListOrder(postIDList)
	// 将postList封装成ApiPostDetail
	//start := (p.Page - 1) * p.Size
	//end := start + p.Size - 1
	apiPostDetailList = make([]*models.ApiPostDetail, 0, len(postList))
	for i := 0; i < len(postList); i++ {
		authorName, err := mysql.GetAuthorNameByUserID(postList[i].AuthorID)
		if err != nil {
			return nil, err
		}
		communityCategoryDetail, err := mysql.GetCommunityDetailByID(postList[i].CommunityID)
		if err != nil {
			return nil, err
		}
		postDetail := &models.ApiPostDetail{
			VoteNum:                 postVoteNumList[i],
			AuthorName:              authorName,
			PostContent:             postList[i],
			CommunityCategoryDetail: communityCategoryDetail,
		}
		apiPostDetailList = append(apiPostDetailList, postDetail)
	}
	return
}
