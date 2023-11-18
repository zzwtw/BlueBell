package models

/*
构造user结构体与数据库中user表相对应，相互绑定
*/

type User struct {
	UserID   int64  `db:"user_id"`
	UserName string `db:"username"`
	PassWord string `db:"password"`
}
