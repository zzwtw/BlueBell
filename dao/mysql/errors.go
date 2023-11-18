package mysql

import "errors"

var (
	ErrorUserExist                    = errors.New("用户已存在")
	ErrorUserNotExist                 = errors.New("用户不存在")
	ErrorLoginUserNameOrPassWordError = errors.New("用户名或密码错误")
)
