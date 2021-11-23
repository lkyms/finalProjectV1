package util

import (
	"github.com/spf13/viper"
)

func GetConfig(str string) (ans string) {
	viper.AddConfigPath("./conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	ans = viper.GetString(str)
	return
}
