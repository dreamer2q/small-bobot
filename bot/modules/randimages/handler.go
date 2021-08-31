package randimages

import (
	"bytes"
	"errors"
	"miraigo-robot/bot"
	"miraigo-robot/config"
	"miraigo-robot/utils"

	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"gorm.io/gorm"
)

type RandImages struct{}

var (
	logger = utils.GetModuleLogger("randimages")
	db     *gorm.DB
)

func init() {
	bot.RegisterModule(RandImages{})
}

func (c RandImages) Module() bot.ModuleInfo {
	return bot.ModuleInfo{
		ID:       "dreamer2q.randimages",
		Instance: c,
	}
}

func (RandImages) Init() {
	logger.Infof("module randimages: started")
	db = config.DB
	_ = db.AutoMigrate(&Image{})
}

func (RandImages) Start(bot *bot.Bot) {
	// register event handler
	bot.OnGroupMessage(onGroupMsg)
}

func (RandImages) Run() {}

func (RandImages) Stop() {
	logger.Infof("module randimages: stopped")
}

var _ bot.Module = RandImages{}

//onGroupMsg handle specified message and make reply
func onGroupMsg(clt *client.QQClient, msg *message.GroupMessage) {
	if msg.ToString() == "求图" {
		rpl := message.NewSendingMessage()
		rpl.Append(message.NewReply(msg))

		reply := func(f func() error) {
			if err := f(); err != nil {
				rpl.Append(message.NewText(err.Error()))
			}
			clt.SendGroupMessage(msg.GroupCode, rpl)
		}

		reply(func() error {
			img, err := GetRandomResult()
			if err != nil {
				logger.Errorf("random images: %v", err)
				return errors.New("数据库错误")
			}
			imgBytes, err := utils.GetBytes(baseUrl + "/" + img.Url)
			if err != nil {
				return errors.New("图片下载失败")
			}
			gm, err := clt.UploadGroupImage(msg.GroupCode, bytes.NewReader(imgBytes))
			if err != nil {
				logger.Errorf("upload group image: %v", err)
				return errors.New("图片上传失败")
			}
			rpl.Append(gm)
			return nil
		})
	}
}
