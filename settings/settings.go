package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigFile("./settings/config.yaml") // 指定配置文件路径
	err = viper.ReadInConfig()                    // 读取配置信息
	if err != nil {                               // 读取配置信息失败
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
		return
	}
	// 监控配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
	})
	return
}
