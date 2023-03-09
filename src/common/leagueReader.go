package common

import (
	"github.com/spf13/viper"
	"log"
	"strings"
)

func init() {
	viper.SetConfigName("league")
	viper.AddConfigPath(".")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		log.Print("配置文件读取错误", err)
	}
}

func Reader(key string) *string {
	if strings.Contains(key, "中超") {
		key = "中超"
	}
	name := viper.GetString(key)
	return &name
}
