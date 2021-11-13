package randsentence

import (
	"fmt"
	"miraigo-robot/bot"
	"miraigo-robot/pkg/hitokoto"
	"miraigo-robot/utils"
	"strings"

	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
)

type m struct{}

var (
	logger = utils.GetModuleLogger("hitokoto")
)

func init() {
	module := &m{}
	bot.RegisterModule(module)
}

func (h *m) Module() bot.ModuleInfo {
	return bot.ModuleInfo{
		ID:       "hdu.hitokoto",
		Instance: h,
	}
}

func (h *m) Init() {
	logger.Infof("module randsentence: started")
}

func (f *m) Start(bot *bot.Bot) {
	bot.OnGroupMessage(func(client *client.QQClient, msg *message.GroupMessage) {
		if msg.ToString() == "一言" {
			r := message.NewSendingMessage().
				Append(message.NewReply(msg))
			hitokoto.SayHi()
			hi, err := hitokoto.GetHitokoto()
			if err != nil {
				r.Append(message.NewText(fmt.Sprintf("一言: %v", err)))
			} else {
				sb := strings.Builder{}
				sb.WriteString(hi.Hitokoto)
				sb.WriteString("\n — ")
				sb.WriteString(hi.From)
				if hi.FromWho != "" {
					sb.WriteString(fmt.Sprintf("(%s)", hi.FromWho))
				}
				r.Append(message.NewText(sb.String()))
			}
			client.SendGroupMessage(msg.GroupCode, r)
		}
	})
}

func (f *m) Run() {}

func (f *m) Stop() {}

var _ bot.Module = &m{}
