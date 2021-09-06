package faq

import (
	"miraigo-robot/bot"
	"miraigo-robot/config"
	"miraigo-robot/utils"
	"regexp"

	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"gorm.io/gorm"
)

type mod struct{}

var (
	logger = utils.GetModuleLogger("faq")

	db *gorm.DB
)

func init() {
	bot.RegisterModule(mod{})
}

func (c mod) Module() bot.ModuleInfo {
	return bot.ModuleInfo{
		ID:       "dreamer2q.faq",
		Instance: c,
	}
}

func (mod) Init() {
	logger.Infof("module faq: started")
	db = config.DB
	db.AutoMigrate(&Faq{}, &FaqLog{})
}

func (mod) Start(bot *bot.Bot) {
	// register event handler
	bot.OnGroupMessage(onGroupMsg)
}

func (mod) Run() {}

func (mod) Stop() {
	logger.Infof("module faq: stopped")
}

var _ bot.Module = mod{}

//onGroupMsg handle specified message and make reply
func onGroupMsg(clt *client.QQClient, msg *message.GroupMessage) {

	textMsg := msg.ToString()
	faqs := make([]Faq, 0)
	rpl := message.NewSendingMessage().Append(message.NewReply(msg))

	err := db.Where(&Faq{GroupID: msg.GroupCode, Enabled: true}).
		Find(&faqs).
		Error
	if err != nil {
		logger.Errorf("db query faqs: %v", err)
		return
	}
	if len(faqs) == 0 {
		return
	}

	for _, v := range faqs {
		faq := v
		rule := faq.Question
		matched, err := regexp.Match(rule, []byte(textMsg))
		if err != nil {
			db.Model(&faq).Update("enabled", false)
			logger.Errorf("faq ruleset: %v", err)
			continue
		}
		if !matched {
			continue
		}
		ans := message.NewText(faq.Answer)
		rpl.Append(ans)
		clt.SendGroupMessage(msg.GroupCode, rpl)
		break
	}
}
