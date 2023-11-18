package snowflake

import (
	sf "github.com/bwmarrin/snowflake"
	"time"
)

var node *sf.Node

func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	//  把时间字符串转换为Time，时区是UTC时区。
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	// 设置起始时间
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	return
}
func GenID() int64 {
	return node.Generate().Int64()
}
