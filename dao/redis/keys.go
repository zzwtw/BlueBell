package redis

const (
	Prefix       = "bluebell:"
	KeyPostTime  = "post:time"  // 帖子创建的时间
	KeyPostScore = "post:score" // 帖子的分数
	KeyPostVote  = "post:vote:" // 帖子的投票
)

func getKey(str string) string {
	return Prefix + str
}
