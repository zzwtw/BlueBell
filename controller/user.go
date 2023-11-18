package controller

import (
	"bluebell/dao/mysql"
	"bluebell/logic"
	"bluebell/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

const CtxUserIdKey = "userID"

func SignupHandler(c *gin.Context) {
	// 1.获取参数和参数校验
	p := new(models.ParamSignUp) // 得到一个注册请求参数结构体指针
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// 判断err类型是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		// 如果不是就返回正常错误
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		// 如果是就返回翻译过后的内容
		ResponseErrorWithMsg(c, CodeInvalidParam, errs.Translate(trans))
		return
	}
	// 2.业务处理 logic层进行业务处理
	if err := logic.Signup(p); err != nil {
		if err == mysql.ErrorUserExist {
			ResponseErrorWithMsg(c, CodeUserExist, mysql.ErrorUserExist.Error())
		}

		return
	}
	// 3.返回响应
	ResponseSuccessWithMsg(c, CodeSuccess, "用户注册成功")
	return
}

func LoginHandler(c *gin.Context) {
	// 1.参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		if errs, ok := err.(validator.ValidationErrors); ok {
			ResponseErrorWithMsg(c, CodeInvalidParam, errs.Translate(trans))
			return
		} else {
			ResponseError(c, CodeInvalidParam)
			return
		}
	}
	// 2.业务处理
	token, err := logic.Login(p)
	fmt.Println(err)
	if err != nil {
		// 存在错误说明 1.用户不存在 2.用户名或密码错误
		if err == mysql.ErrorUserNotExist {
			ResponseErrorWithMsg(c, CodeUserNotExist, mysql.ErrorUserNotExist.Error())
		} else if err == mysql.ErrorLoginUserNameOrPassWordError {
			ResponseErrorWithMsg(c, CodeUserOrPassWordError, mysql.ErrorLoginUserNameOrPassWordError.Error())
		}
		// 存入token相关的错误到日志中
		zap.L().Error("logic.Login another err", zap.Error(err))
		return
	}
	// 3.返回响应
	ResponseSuccessWithMsg(c, CodeSuccess, token)
	return
}
