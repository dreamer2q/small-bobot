package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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
	GlobalConfig.AddConfigPath("/etc/smbot")

	err := GlobalConfig.ReadInConfig()
	if err != nil {
		logrus.WithField("config", "GlobalConfig").WithError(err).Panicf("unable to read global config")
	}

	initDB()
}

//ReadDeviceJson read device.json file
func ReadDeviceJson() []byte {
	const device = `{
		"display": "MIRAI.415395.001",
		"product": "mirai",
		"device": "mirai",
		"board": "mirai",
		"model": "mirai",
		"finger_print": "mamoe/mirai/mirai:10/MIRAI.200122.001/0223216:user/release-keys",
		"boot_id": "9539d34a-bec8-16f6-8df7-5eb3158cf6cf",
		"proc_version": "Linux version 3.0.31-7gBfR6A3 (android-build@xxx.xxx.xxx.xxx.com)",
		"imei": "762641695676147"
	}`
	// return utils.ReadFile("./config/device.json")
	return []byte(device)
}
