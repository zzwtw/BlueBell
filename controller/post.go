package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"bluebell/pkg/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"strconv"
)

/*
帖子相关
*/
const OrderByTime = "time"

// CreatePostHandler 处理用户新增帖子
func CreatePostHandler(c *gin.Context) {
	// 1.参数校验
	p := new(models.PostContent)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("createPostHandler invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseSuccessWithMsg(c, CodeInvalidParam, errs.Translate(trans))
		return
	}
	// 2.业务层处理
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
	}
	p.ID = snowflake.GenID()
	p.AuthorID = userID
	if err = logic.CreatePost(p); err != nil {
		zap.L().Error("logic CreatePost error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3.返回响应
	ResponseSuccess(c, CodeSuccess)
	return
}

func PostDetailByIDHandler(c *gin.Context) {
	postIDStr := c.Param("id")
	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		zap.L().Error("postId invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	apiPostDetail := new(models.ApiPostDetail)
	apiPostDetail, err = logic.GetPostDetailByID(postID)
	if err != nil {
		zap.L().Error("logic GetPostDetailByID error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, apiPostDetail)
	return
}

func PostListHandler(c *gin.Context) {
	// 获取当前page以及size
	// get参数用Query获取
	pageStr := c.Query("page")
	sizeStr := c.Query("size")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		size = 10
	}
	// logic业务处理
	postContentList, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic getPostList error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, postContentList)
	return

}

// 接受参数，判断是按时间顺序返回还是分数大小返回

func PostListOrderHandler(c *gin.Context) {
	size, err := strconv.Atoi(c.Query("size"))
	page, err := strconv.Atoi(c.Query("page"))
	p := &models.ParamPost{
		Page:  page,
		Size:  size,
		Order: c.Query("order"),
	}
	if err != nil {
		zap.L().Error("post list order invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	apiPostDetailList, err := logic.GetPostIDListOrderByTime(p)
	if err != nil {
		zap.L().Error("logic GetPostIDListOrderByTime error ", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, apiPostDetailList)
	return
}
