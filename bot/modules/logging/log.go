package logging

import (
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"gorm.io/gorm"
	"miraigo-robot/bot"
	"miraigo-robot/config"
	"miraigo-robot/utils"
)

func init() {
	instance = &logging{}
	bot.RegisterModule(instance)
}

var (
	db *gorm.DB
)

type logging struct{}

func (m *logging) Module() bot.ModuleInfo {
	return bot.ModuleInfo{
		ID:       "internal.logging",
		Instance: instance,
	}
}

func (m *logging) Init() {
	// 初始化过程
	// 在此处可以进行 Module 的初始化配置
	// 如配置读取
	db = config.DB
	_ = db.AutoMigrate(
		&PrivateMessage{},
		&GroupMessage{},
		&Group{},
		&Friend{})
}

func (m *logging) Start(b *bot.Bot) {
	groups, _ := b.GetGroupList()
	for _, group := range groups {
		db.Save(&Group{
			ID:              group.Code,
			Name:            group.Name,
			Memo:            group.Memo,
			OwnerID:         group.OwnerUin,
			GroupCreateTime: group.GroupCreateTime,
			GroupLevel:      group.GroupLevel,
			MemberCount:     group.MemberCount,
			MaxMemberCount:  group.MaxMemberCount,
		})
	}
	friends, _ := b.GetFriendList()
	for _, friend := range friends.List {
		info, _ := b.GetSummaryInfo(friend.Uin)
		db.Save(&Friend{
			ID:        friend.Uin,
			Remark:    friend.Remark,
			Sex:       info.Sex,
			Age:       info.Age,
			Nickname:  info.Nickname,
			Level:     info.Level,
			City:      info.City,
			Sign:      info.Sign,
			Mobile:    info.Mobile,
			LoginDays: info.LoginDays,
			Qid:       info.Qid,
		})
	}
	// 注册服务函数部分
	registerLog(b)
}

func (m *logging) Run() {
	// 此函数会新开携程进行调用
	// ```go
	// 		go exampleModule.Run()
	// ```

	// 可以利用此部分进行后台操作
	// 如http服务器等等
}

func (m *logging) Stop() {
	// 结束部分
	// 一般调用此函数时，程序接收到 os.Interrupt 信号
	// 即将退出
	// 在此处应该释放相应的资源或者对状态进行保存
	// 取消订阅事件
}

var instance *logging

var logger = utils.GetModuleLogger("internal.logging")

func logGroupMessage(msg *message.GroupMessage) {
	db.Create(&GroupMessage{
		MessageBase: MessageBase{
			Time:       msg.Time,
			MessageId:  msg.Id,
			InternalId: msg.InternalId,
			SenderID:   msg.Sender.Uin,
			Message:    msg.ToString(),
		},
		GroupCode: msg.GroupCode,
	})
	logger.
		WithField("from", "GroupMessage").
		WithField("MessageID", msg.Id).
		WithField("MessageIID", msg.InternalId).
		WithField("GroupCode", msg.GroupCode).
		WithField("SenderID", msg.Sender.Uin).
		Info(msg.ToString())

}

func logPrivateMessage(msg *message.PrivateMessage) {
	db.Create(&PrivateMessage{
		MessageBase: MessageBase{
			Time:       msg.Time,
			MessageId:  msg.Id,
			InternalId: msg.InternalId,
			SenderID:   msg.Sender.Uin,
			Message:    msg.ToString(),
		},
		Self:   msg.Self,
		Target: msg.Target,
	})
	logger.
		WithField("from", "PrivateMessage").
		WithField("MessageID", msg.Id).
		WithField("MessageIID", msg.InternalId).
		WithField("SenderID", msg.Sender.Uin).
		WithField("Target", msg.Target).
		Info(msg.ToString())
}

func logFriendMessageRecallEvent(event *client.FriendMessageRecalledEvent) {
	logger.
		WithField("from", "FriendsMessageRecall").
		WithField("MessageID", event.MessageId).
		WithField("SenderID", event.FriendUin).
		Info("friend message recall")
}

func logGroupMessageRecallEvent(event *client.GroupMessageRecalledEvent) {
	logger.
		WithField("from", "GroupMessageRecall").
		WithField("MessageID", event.MessageId).
		WithField("GroupCode", event.GroupCode).
		WithField("SenderID", event.AuthorUin).
		WithField("OperatorID", event.OperatorUin).
		Info("group message recall")
}

func logGroupMuteEvent(event *client.GroupMuteEvent) {
	logger.
		WithField("from", "GroupMute").
		WithField("GroupCode", event.GroupCode).
		WithField("OperatorID", event.OperatorUin).
		WithField("TargetID", event.TargetUin).
		WithField("MuteTime", event.Time).
		Info("group mute")
}

func logDisconnect(event *client.ClientDisconnectedEvent) {
	logger.
		WithField("from", "Disconnected").
		WithField("reason", event.Message).
		Warn("bot disconnected")
}

func registerLog(b *bot.Bot) {
	b.OnGroupMessageRecalled(func(qqClient *client.QQClient, event *client.GroupMessageRecalledEvent) {
		logGroupMessageRecallEvent(event)
	})

	b.OnGroupMessage(func(qqClient *client.QQClient, groupMessage *message.GroupMessage) {
		logGroupMessage(groupMessage)
	})

	b.OnGroupMuted(func(qqClient *client.QQClient, event *client.GroupMuteEvent) {
		logGroupMuteEvent(event)
	})

	b.OnPrivateMessage(func(qqClient *client.QQClient, privateMessage *message.PrivateMessage) {
		logPrivateMessage(privateMessage)
	})

	b.OnFriendMessageRecalled(func(qqClient *client.QQClient, event *client.FriendMessageRecalledEvent) {
		logFriendMessageRecallEvent(event)
	})

	b.OnDisconnected(func(qqClient *client.QQClient, event *client.ClientDisconnectedEvent) {
		logDisconnect(event)
	})
}
