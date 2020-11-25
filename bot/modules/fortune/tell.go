package fortune

import (
	"fmt"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	telling "github.com/dreamer2q/fortune_telling"
	"log"
	"miraigo-robot/bot"
	"strings"
	"sync"
	"text/template"
)

type Fortune struct{}

func (f Fortune) Module() bot.ModuleInfo {
	return bot.ModuleInfo{
		ID:       "dreamer2q.fortune",
		Instance: f,
	}
}

func (f Fortune) Init() {
	tell, _ := telling.Ask("test")
	log.Printf("telling: %v", tell)
}

func (f Fortune) Serve(bot *bot.Bot) {
	bot.OnGroupMessage(func(client *client.QQClient, msg *message.GroupMessage) {
		if msg.ToString() == "求签" {
			tel, err := telling.Ask(fmt.Sprintf("%v", msg.Sender.Uin))
			if err != nil {
				client.SendGroupMessage(msg.GroupCode, message.NewSendingMessage().Append(message.NewText("重复求签")))
				return
			}
			tpl := template.Must(template.New("sign").
				Funcs(template.FuncMap{
					"parseJi": tel.String,
				}).
				Parse(
					`
🌓运势：{{.Level}}
🌟指数：{{.Level | parseJi}}
📗签文：{{.Content}}
📝解签：{{.Detail1}}
☯说签：{{.Detail2}}`,
				))
			sb := &strings.Builder{}
			err = tpl.Execute(sb, &tel)

			m := message.NewSendingMessage()
			m.Append(message.NewAt(msg.Sender.Uin))
			m.Append(message.NewText("\n"))
			m.Append(message.NewText(sb.String()))
			client.SendGroupMessage(msg.GroupCode, m)
		}
	})
}

func (f Fortune) Start(bot *bot.Bot) {
	panic("implement me")
}

func (f Fortune) Stop(bot *bot.Bot, wg *sync.WaitGroup) {
	panic("implement me")
}

var _ bot.Module = Fortune{}
