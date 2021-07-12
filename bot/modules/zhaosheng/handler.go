package zhaosheng

import (
	"fmt"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"miraigo-robot/bot"
	"miraigo-robot/utils"
	"strings"
)

type zhaosheng struct{}

var (
	logger = utils.GetModuleLogger("zhaosheng")
)

func init() {
	bot.RegisterModule(zhaosheng{})
}

func (z zhaosheng) Module() bot.ModuleInfo {
	return bot.ModuleInfo{
		ID:       "hdu.zhaosheng",
		Instance: z,
	}
}

func (z zhaosheng) Init() {
	logger.Infof("hdu.zhaosheng started")
}

func (z zhaosheng) Start(bot *bot.Bot) {
	bot.OnNewFriendRequest(func(client *client.QQClient, request *client.NewFriendRequest) {
		request.Accept()
	})
	bot.OnNewFriendAdded(func(client *client.QQClient, event *client.NewFriendEvent) {
		client.SendPrivateMessage(event.Friend.Uin, message.NewSendingMessage().
			Append(message.NewText("HDU导航 | hduer.cn，为你服务")))
	})
	bot.OnTempMessage(func(client *client.QQClient, event *client.TempMessageEvent) {
		_, _ = event.Session.SendMessage(message.NewSendingMessage().
			Append(message.NewText("有事吗？请先加好友。")))
	})
	bot.OnPrivateMessageF(func(msg *message.PrivateMessage) bool {
		str := msg.ToString()
		// zscx ksh xm
		return strings.HasPrefix(str, "zs")
	}, func(client *client.QQClient, p *message.PrivateMessage) {
		params := strings.Split(p.ToString(), " ")
		reply := func(s string) {
			logger.Infof("send to %v: %v", p.Sender.Uin, s)
			client.SendPrivateMessage(p.Sender.Uin, message.NewSendingMessage().Append(message.NewText(s)))
		}
		res := func() string {
			const (
				help      = "回复 zshelp 获取帮助"
				helpUsage = "格式：zscx 准考号 姓名"
			)
			cmd := params[0]
			params = params[1:]
			switch cmd {
			case "zshelp":
				return "招生查询小助手：\n" +
					"录取查询  " + helpUsage +
					"\n欢迎你的测试与反馈"
			case "zscx":
				if len(params) != 2 {
					return helpUsage
				}
				res, err := GetQueryResult(params[0], params[1])
				if err != nil {
					return fmt.Sprintf("出错了：%v", err)
				}
				return res
			}
			return help
		}
		reply(res())
	})
}

func (z zhaosheng) Run() {}

func (z zhaosheng) Stop() {}

var _ bot.Module = zhaosheng{}
