package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"miraigo-robot/utils"
)

type Config struct {
	*viper.Viper
}

// GlobalConfig 默认全局配置
var GlobalConfig *Config

//init 使用 config.yml 初始化全局配置
//请确保在使用 GlobalConfig 之前进行初始化
func init() {
	GlobalConfig = &Config{
		viper.New(),
	}
	GlobalConfig.SetConfigName("config")
	GlobalConfig.SetConfigType("yaml")
	GlobalConfig.AddConfigPath(".")
	GlobalConfig.AddConfigPath("./config")

	err := GlobalConfig.ReadInConfig()
	if err != nil {
		logrus.WithField("config", "GlobalConfig").WithError(err).Panicf("unable to read global config")
	}

	initDB()
}

//ReadDeviceJson read device.json file
func ReadDeviceJson() []byte {
	return utils.ReadFile("./config/device.json")
}
