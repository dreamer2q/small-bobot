package purify

import (
	"miraigo-robot/bot"
	"miraigo-robot/utils"
	"regexp"

	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
)

type module struct{}

var (
	logger = utils.GetModuleLogger("purify")
)

func init() {
	bot.RegisterModule(module{})
}

func (c module) Module() bot.ModuleInfo {
	return bot.ModuleInfo{
		ID:       "dreamer2q.purify",
		Instance: c,
	}
}

func (module) Init() {
	logger.Infof("module purify: started")
}

func (module) Start(bot *bot.Bot) {
	// register event handler
	bot.OnGroupMessage(onGroupMsg)
}

func (module) Run() {}

func (module) Stop() {
	logger.Infof("module purify: stopped")
}

var _ bot.Module = module{}

var purifyRule = regexp.MustCompile(`傻逼`)

//onGroupMsg handle specified message and make reply
func onGroupMsg(clt *client.QQClient, msg *message.GroupMessage) {
	const serveGroup = 924962654
	if msg.GroupCode != serveGroup {
		return
	}

	mi, err := clt.GetMemberInfo(msg.GroupCode, msg.Sender.Uin)
	if err != nil {
		logger.Errorf("get group member info: %v", err)
		return
	}

	if utils.IsAt(msg, clt.Uin) {
		err := clt.RecallGroupMessage(msg.GroupCode, msg.Id, msg.InternalId)
		if err != nil {
			logger.Errorf("recall group message: %v", err)
		}
		mi.Mute(300)
		return
	}

	textMsg := utils.GetGroupTextMsg(msg)
	needPurify := purifyRule.MatchString(textMsg)
	if needPurify {
		err := clt.RecallGroupMessage(msg.GroupCode, msg.Id, msg.InternalId)
		if err != nil {
			logger.Errorf("recall group message: %v", err)
		}
		mi.Mute(30)
	}
}
