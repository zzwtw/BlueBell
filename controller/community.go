package controller

import (
	"bluebell/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

/*
社区分类相关
*/

func CommunityListHandler(c *gin.Context) {
	communityList, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("获取社区帖子分类失败", zap.Error(err))
		ResponseErrorWithMsg(c, CodeServerBusy, err)
	}
	ResponseSuccess(c, communityList)
}

func CommunityDetailByIDHandler(c *gin.Context) {
	// 获取所需要查询的社区分类的id
	communityIDStr := c.Param("id")
	communityID, err := strconv.ParseInt(communityIDStr, 10, 64)
	if err != nil {
		zap.L().Error("communityID invalid param", zap.Error(err))
	}
	communityDetail, err := logic.GetCommunityDetailByID(communityID)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetailByID error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}
	ResponseSuccess(c, communityDetail)
}
