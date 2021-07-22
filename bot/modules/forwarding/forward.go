package forwarding

import (
	"fmt"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"miraigo-robot/bot"
	"miraigo-robot/config"
	"regexp"
)

func init() {
	module := &forward{}
	bot.RegisterModule(module)
}

type forward struct {
	toGroup int64
}

func (f *forward) Module() bot.ModuleInfo {
	return bot.ModuleInfo{
		ID:       "hdu.forwarding",
		Instance: f,
	}
}

func (f *forward) Init() {
	c := config.GlobalConfig
	f.toGroup = c.GetInt64("module.forwarding.toGroup")
}

func (f *forward) Start(bot *bot.Bot) {
	var regx = regexp.MustCompile(`办卡|移动|电信|宽带|闪讯|拨号`)
	bot.OnGroupMessage(func(client *client.QQClient, msg *message.GroupMessage) {
		var txt string
		for _, elem := range msg.Elements {
			if elem.Type() == message.Text {
				txt += elem.(*message.TextElement).Content
			}
		}
		if regx.MatchString(txt) {
			var elems = []message.IMessageElement{
				message.NewText(fmt.Sprintf("《%s》%s（%v）触发关键词\n",
					msg.GroupName, msg.Sender.DisplayName(), msg.Sender.Uin))}
			elems = append(elems, msg.Elements...)
			client.SendGroupMessage(f.toGroup, &message.SendingMessage{Elements: elems})
		}
	})
	bot.OnGroupMemberJoined(func(client *client.QQClient, event *client.MemberJoinGroupEvent) {
		client.SendGroupMessage(f.toGroup, message.NewSendingMessage().
			Append(message.NewText(fmt.Sprintf("%s（%v）加入《%s》",
				event.Member.DisplayName(), event.Member.Uin, event.Group.Name))))
	})
	bot.OnGroupMemberLeaved(func(client *client.QQClient, event *client.MemberLeaveGroupEvent) {
		client.SendGroupMessage(f.toGroup, message.NewSendingMessage().
			Append(message.NewText(fmt.Sprintf("%s（%v）退出《%s》",
				event.Member.DisplayName(), event.Member.Uin, event.Group.Name))))
	})
}

func (f *forward) Run() {}

func (f *forward) Stop() {}

var _ bot.Module = &forward{}
