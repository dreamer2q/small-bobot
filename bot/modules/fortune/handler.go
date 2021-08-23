package fortune

import (
	"fmt"
	"miraigo-robot/bot"
	"strings"
	"text/template"

	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	telling "github.com/dreamer2q/fortune_telling"
)

var (
	signTmpl = template.Must(template.New("sign-tmpl").
		Funcs(template.FuncMap{
			"Stars": telling.LevelStars,
		}).
		Parse(
			`
🌓运势：{{.Level}}
🌟指数：{{.Level | Stars}}
📗签文：{{.Content}}
📝解签：{{.Detail1}}
☯说签：{{.Detail2}}`,
		))
)

func registryEvent(bot *bot.Bot) {
	bot.OnGroupInvited(onGroupInvite)
	bot.OnGroupMessage(onGroupMsg)
}

//onGroupInvite accept admin invitation
func onGroupInvite(clt *client.QQClient, req *client.GroupInvitedRequest) {
	if req.InvitorUin == admin {
		req.Accept()
	}
}

//onGroupMsg handle specified message and make reply
func onGroupMsg(clt *client.QQClient, msg *message.GroupMessage) {
	if msg.ToString() == "求签" {
		tell, err := telling.Ask(fmt.Sprintf("%v", msg.Sender.Uin))
		rpl := message.NewSendingMessage()
		rpl.Append(message.NewAt(msg.Sender.Uin))
		rpl.Append(message.NewText("\n"))

		if err != nil {
			rpl.Append(message.NewText("你已经求过签了，请明天再来吧"))
		} else {
			sb := &strings.Builder{}
			err = signTmpl.Execute(sb, &tell)
			rpl.Append(message.NewText(sb.String()))
		}
		clt.SendGroupMessage(msg.GroupCode, rpl)
	}
}
