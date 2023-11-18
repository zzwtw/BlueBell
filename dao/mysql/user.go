package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"encoding/hex"
	"go.uber.org/zap"
)

/*
用于用户表的一些校验
*/

const secret = "zhouweitong"

// CheckUserExist 判断用户是否存在
func CheckUserExist(user *models.ParamSignUp) error {
	sqlStr := "select count(user_id) from user where username = ?"
	var count int
	if err := db.Get(&count, sqlStr, user.Username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return nil
}

// InsertUser 向数据库中插入一条用户数据
func InsertUser(user *models.User) error {
	sqlStr := "insert into user(user_id,username,password) values(?,?,?)"
	user.PassWord = encryptPassword(user.PassWord)
	if _, err := db.Exec(sqlStr, user.UserID, user.UserName, user.PassWord); err != nil {
		return err
	}
	return nil
}

// 对密码进行加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

// CheckLoginUserExist 查询登录用户是否存在
func CheckLoginUserExist(user *models.ParamLogin) error {
	sqlStr := "select count(user_id) from user where username = ?"
	var count int
	if err := db.Get(&count, sqlStr, user.Username); err != nil {
		zap.L().Error("验证用户名是否存在出错", zap.Error(err))
		return err
	}
	if count > 0 {
		return nil
	}
	return ErrorUserNotExist
}

// CheckLoginUserPassword 验证用户名或密码是否正确
func CheckLoginUserPassword(user *models.User) (err error) {
	user.PassWord = encryptPassword(user.PassWord)
	sqlStr := "select user_id,username,password from user where username = ?"
	// 不能是指针类型
	var u models.User
	if err = db.Get(&u, sqlStr, user.UserName); err != nil {
		zap.L().Error("验证用户名密码是否正确出错", zap.Error(err))
		return err
	}
	if user.PassWord == u.PassWord {
		user.UserID = u.UserID
		return nil
	}
	return ErrorLoginUserNameOrPassWordError
}

func GetAuthorNameByUserID(userID int64) (string, error) {
	sqlStr := "select username from user where user_id = ?"
	var authorName string
	err := db.Get(&authorName, sqlStr, userID)

	return authorName, err
}
