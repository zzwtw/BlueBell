package redis

import "bluebell/models"

const OrderByScore = "score"

func GetPostIDListOrderByTimeOrScore(p *models.ParamPost) ([]string, error) {
	// 根据时间或分数顺序返回post_id
	key := getKey(KeyPostTime)
	if p.Order == OrderByScore {
		key = getKey(KeyPostScore)
	}
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	postIDList, err := rdb.ZRevRange(key, int64(start), int64(end)).Result()
	return postIDList, err

}

// GetPostVoteNumByPostID 查询每个帖子的投票数
func GetPostVoteNumByPostID(postIDList []string) (postVoteNumList []int64, err error) {
	for _, post_id := range postIDList {
		result, err := rdb.ZCount(getKey(KeyPostVote+post_id), "1", "1").Result()
		if err != nil {
			return nil, err
		}
		postVoteNumList = append(postVoteNumList, result)
	}
	return
}
