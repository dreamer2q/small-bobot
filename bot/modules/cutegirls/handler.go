package cutegirls

import (
	"bytes"
	"miraigo-robot/bot"
	"miraigo-robot/utils"

	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
)

type Cute struct{}

var (
	logger = utils.GetModuleLogger("cutegirls")
)

func init() {
	bot.RegisterModule(Cute{})
}

func (c Cute) Module() bot.ModuleInfo {
	return bot.ModuleInfo{
		ID:       "dreamer2q.cutegirls",
		Instance: c,
	}
}

func (Cute) Init() {
	logger.Infof("module cutegirls: started")
}

func (Cute) Start(bot *bot.Bot) {
	// register event handler
	bot.OnGroupMessage(onGroupMsg)
}

func (Cute) Run() {}

func (Cute) Stop() {
	logger.Infof("module cutegirls: stopped")
}

var _ bot.Module = Cute{}

//onGroupMsg handle specified message and make reply
func onGroupMsg(clt *client.QQClient, msg *message.GroupMessage) {
	if msg.ToString() == "求图" {
		rpl := message.NewSendingMessage()
		rpl.Append(message.NewReply(msg))

		girl, err := WantCuteGirl()
		if err != nil {
			logger.Errorf("cutegirls: %v", err)
			rpl.Append(message.NewText("恭喜: 求到了一张空气图"))
		} else {
			gm, err := clt.UploadGroupImage(msg.GroupCode, bytes.NewReader(girl.Body))
			if err != nil {
				logger.Warnf("upload group image: %v", err)
				rpl.Append(message.NewText("恭喜: 你的图片被怪兽吃掉了"))
			} else {
				rpl.Append(gm)
			}
		}

		clt.SendGroupMessage(msg.GroupCode, rpl)
	}
}
