package app

import (
	"miraigo-robot/bot"
	"miraigo-robot/config"
	"os"
	"os/signal"

	_ "miraigo-robot/bot/modules/fortune"
	_ "miraigo-robot/bot/modules/logging"
	_ "miraigo-robot/bot/modules/zhaosheng"
)

func Init() {
	// 快速初始化
	conf := config.GlobalConfig
	bot.Init(bot.Config{
		Account:  conf.GetInt64("bot.account"),
		Password: conf.GetString("bot.password"),
		Device:   config.ReadDeviceJson(),
	})
	// 使用协议
	// 不同协议可能会有部分功能无法使用
	// 在登陆前切换协议
	bot.UseProtocol(bot.AndroidPhone)
	// 登录
	bot.Login()
	// 初始化 Modules
	bot.StartService()
	// 刷新好友列表，群列表
	bot.RefreshList()
	// wait
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill)
	<-ch
	bot.Stop()
}
