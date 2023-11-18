package redis

import (
	"bluebell/settings"
	"fmt"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func Init(config settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			//viper.GetString("redis.host"),
			//viper.GetInt("redis.port"),
			config.Host,
			config.Port,
		),
		Password: config.PassWord,
		DB:       config.DB,
		PoolSize: config.PoolSize,
	})
	_, err = rdb.Ping().Result()
	return
}

func Close() {
	_ = rdb.Close()
}
