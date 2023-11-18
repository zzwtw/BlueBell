package logic

import (
	"bluebell/dao/redis"
	"go.uber.org/zap"
	"strconv"
)

func PostVote(postID int64, direction int, userID int64) error {
	// 判断该帖子是否过了投票时间
	if err := redis.CheckVoteTime(postID); err != nil {
		zap.L().Error("redis checkVoteTime error", zap.Error(err))
		return err
	}
	// 拿到userID给指定的帖子(postID)投票
	userIDStr := strconv.Itoa(int(userID))
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return err
	}
	if err = redis.UserPostVote(userID, postID, direction); err != nil {
		zap.L().Error("redis userPostVote error", zap.Error(err))
		return err
	}
	return err
}
