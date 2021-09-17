package welcome

import (
	"miraigo-robot/bot"
	"miraigo-robot/config"
	"miraigo-robot/utils"

	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"gorm.io/gorm"
)

var (
	logger = utils.GetModuleLogger("welcome")

	db *gorm.DB
)

func init() {
	module := &moduler{}
	bot.RegisterModule(module)
}

type moduler struct{}

func (m *moduler) Module() bot.ModuleInfo {
	return bot.ModuleInfo{
		ID:       "hdu.welcome",
		Instance: m,
	}
}

func (m *moduler) Init() {
	logger.Infof("module welcome: started")
	db = config.DB
	db.AutoMigrate(&Welcome{})
}

func (m *moduler) Start(bot *bot.Bot) {

	bot.OnGroupMemberJoined(func(client *client.QQClient, event *client.MemberJoinGroupEvent) {

		newbie := event.Member
		rpl := message.NewSendingMessage().
			Append(message.NewAt(newbie.Uin))

		welcome := Welcome{}

		err := db.Where(&Welcome{GroupID: event.Group.Code, Enabled: true}).
			First(&welcome).Error
		if err != nil {
			logger.Errorf("db query welcome: %v", err)
			return
		}
		if welcome.ID == 0 {
			return
		}
		rpl.Append(message.NewText(welcome.Welcome))

		client.SendGroupMessage(event.Group.Code, rpl)
	})

	// bot.OnGroupMemberLeaved(func(client *client.QQClient, event *client.MemberLeaveGroupEvent) {})
}

func (m *moduler) Run() {}

func (m *moduler) Stop() {
	logger.Infof("module welcome: stopped")
}

var _ bot.Module = &moduler{}
