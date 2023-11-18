package models

type Vote struct {
	PostID    int64 `json:"post_id,string" binding:"required"`
	Direction int   `json:"direction"`
}
