package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
)

func Signup(p *models.ParamSignUp) error {
	// 1.校验用户是否存在
	if err := mysql.CheckUserExist(p); err != nil {
		return err
	}
	// 2.生成UID
	UID := snowflake.GenID()
	// 生成一个用户实例
	user := &models.User{
		UserID:   UID,
		UserName: p.Username,
		PassWord: p.Password,
	}
	// 3.保存到数据库
	if err := mysql.InsertUser(user); err != nil {
		return err
	}
	return nil
}

// Login 登录业务处理
func Login(p *models.ParamLogin) (string, error) {
	// 构造一个用户表结构体
	user := &models.User{
		UserName: p.Username,
		PassWord: p.Password,
	}
	// 1.判断该用户是否存在，如果不存在则返回错误
	if err := mysql.CheckLoginUserExist(p); err != nil {
		// 存在错误说明用户不存在
		return "", err
	}
	// 2.验证密码是否正确
	if err := mysql.CheckLoginUserPassword(user); err != nil {
		// 存在错误说明用户名或密码错误
		return "", err
	}
	// 拿到userid和username生成token返回给controller
	token, err := jwt.GenToken(user.UserID, user.UserName)
	if err != nil {
		return "", err
	}
	return token, nil
}
