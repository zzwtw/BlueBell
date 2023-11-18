package mysql

import (
	"bluebell/models"
	"github.com/jmoiron/sqlx"
	"strings"
)

func InsetNewPostContent(postContent *models.PostContent) (err error) {
	sqlStr := "insert into post (post_id,title,content,author_id,community_id) values(?,?,?,?,?)"
	_, err = db.Exec(sqlStr, postContent.ID, postContent.Title, postContent.Content, postContent.AuthorID, postContent.CommunityID)
	return
}

func GetPostDetailByID(postID int64) (postContentById *models.PostContent, err error) {
	sqlStr := "select post_id, title, content, author_id, community_id, create_time from post where post_id = ?"
	postContentById = new(models.PostContent)
	err = db.Get(postContentById, sqlStr, postID)
	return
}

func GetPostList(page, size int) (postContentList []*models.PostContent, err error) {
	sqlStr := "select post_id, title, content, author_id, community_id,create_time from post limit ?,?"
	postContentList = make([]*models.PostContent, 0, size)
	// 从第(page-1)*size行开始查询，查询size条数据
	err = db.Select(&postContentList, sqlStr, (page-1)*size, size)
	return
}

func GetPostListOrder(postIDList []string) (postList []*models.PostContent, err error) {
	postList = []*models.PostContent{}
	sqlStr := "select post_id, title, content, author_id, community_id, create_time from post where post_id in (?) order by FIND_IN_SET (post_id,?)"
	query, args, err := sqlx.In(sqlStr, postIDList, strings.Join(postIDList, ","))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)
	err = db.Select(&postList, query, args...)
	if err != nil {
		return nil, err
	}
	return
}
