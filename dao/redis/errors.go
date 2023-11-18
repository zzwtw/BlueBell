package redis

import "errors"

var (
	ErrorVoteTime = errors.New("帖子投票时间已过")
)
