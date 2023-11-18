package settings

// 初始化viper
import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name        string `mapstructure:"name"`
	Mode        string `mapstructure:"mode"`
	Port        int    `mapstructure:"port"`
	Version     string `mapstructure:"version"`
	StartTime   string `mapstructure:"start_time"`
	MachineID   int64  `mapstructure:"machine_id"`
	LogConfig   `mapstructure:"log"`
	MySQLConfig `mapstructure:"mysql"`
	RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackUps int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	PassWord     string `mapstructure:"password"`
	DBName       string `mapstructure:"dbname"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	PassWord string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init() (err error) {
	// 指定配置文件
	// 方式1：直接指定配置文件路径（相对路径或者绝对路径）
	// 相对路径：相对执行的可执行文件的相对路径
	viper.SetConfigFile("./conf/config.yaml")
	// 指定在什么路径下查找配置文件(相对路径)
	//viper.AddConfigPath("./conf")
	// 读取配置文件信息
	if err = viper.ReadInConfig(); err != nil {
		return
	}
	// 反序列化到Conf中
	if err = viper.Unmarshal(Conf); err != nil {
		return err
	}
	// 持续监视配置文件是否发生变化
	viper.WatchConfig()
	// 配置文件发生改变执行的回调函数
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件被修改了")
		// 重新反序列化到Conf中
		if err = viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		}
	})
	return nil
}
