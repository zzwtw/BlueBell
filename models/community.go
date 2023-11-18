package models

type CommunityCategoryList struct {
	ID   int64  `json:"community_id" db:"community_id"`
	Name string `json:"community_name" db:"community_name"`
}

type CommunityCategoryDetail struct {
	ID           int    `json:"community_id" db:"community_id"`
	Name         string `json:"community_name" db:"community_name"`
	Introduction string `json:"community_introduction" db:"introduction"`
}
