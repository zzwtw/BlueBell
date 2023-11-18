package models

import "time"

type PostContent struct {
	ID          int64     `json:"post_id" db:"post_id"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	AuthorID    int64     `json:"author_id" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

type ApiPostDetail struct {
	AuthorName string `json:"author_name"`
	VoteNum    int64  `json:"vote_num"`
	*PostContent
	*CommunityCategoryDetail
}
