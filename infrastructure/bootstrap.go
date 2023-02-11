package infrastructure

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go-chatbot/config"
)

var AppConfig = new(config.OpenaiKeyConfig)

func InitConfig()  {
	v := viper.New()
	v.SetConfigName("application")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	// 文件路径如何设置
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	// 配置文件对象 需要配置为全局变量
	if err := v.Unmarshal(AppConfig); err != nil {
		panic(err)
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		_ = v.ReadInConfig()
		_ = v.Unmarshal(AppConfig)
	})
}