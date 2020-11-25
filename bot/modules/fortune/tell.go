package fortune

import (
	"github.com/Mrs4s/MiraiGo/message"
	"miraigo-robot/bot"
	"miraigo-robot/config"
	"miraigo-robot/utils"
	"sync"
)

type Fortune struct{}

var (
	logger = utils.GetModuleLogger("fortune")
	admin  int64 //uint for admin
)

func init() {
	bot.RegisterModule(Fortune{})
}

func (f Fortune) Module() bot.ModuleInfo {
	return bot.ModuleInfo{
		ID:       "dreamer2q.fortune",
		Instance: f,
	}
}

func (f Fortune) Init() {
	admin = config.GlobalConfig.GetInt64("bot.admin")
	logger.Infof("module fortune: started")
}

func (f Fortune) Serve(bot *bot.Bot) {
	registryEvent(bot)
}

func (f Fortune) Start(bot *bot.Bot) {
	//notify admin
	logger.Infof("notify admin: %d", admin)
	bot.SendPrivateMessage(admin, message.NewSendingMessage().Append(message.NewText("fortune started")))
}

func (f Fortune) Stop(bot *bot.Bot, wg *sync.WaitGroup) {
	logger.Infof("module fortune: stopped")
}

var _ bot.Module = Fortune{}
