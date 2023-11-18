package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func PostVoteHandler(c *gin.Context) {
	// 参数校验
	p := new(models.Vote)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("c shouldBindJson error ", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			ResponseErrorWithMsg(c, CodeInvalidParam, errs.Translate(trans))
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	// 判断该帖子是否过了投票时间
	// 拿到userID给指定的帖子(postID)投票，zset中记录帖子投票数的有序集合为帖子id（postid）中的元素名称即为userID,分数为投票数
	// 更改该帖子的分数
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("getCurrentUserID error ", zap.Error(err))
		ResponseError(c, CodeNeedLogin)
		return
	}
	if err = logic.PostVote(p.PostID, p.Direction, userID); err != nil {
		zap.L().Error("logic PostVote error ", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, CodeVoteSuccess)
	return
}
