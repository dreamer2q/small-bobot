package netease

import (
	"bytes"
	"errors"
	"fmt"
	"miraigo-robot/bot"
	"miraigo-robot/utils"
	"strings"

	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/gabriel-vasile/mimetype"
)

type Song struct{}

var (
	logger           = utils.GetModuleLogger("netease")
	lawfulAudioTypes = [...]string{
		"audio/mpeg", "audio/flac", "audio/midi", "audio/ogg",
		"audio/ape", "audio/amr", "audio/wav", "audio/aiff",
		"audio/mp4", "audio/aac", "audio/x-m4a",
	}
)

func init() {
	bot.RegisterModule(Song{})
}

func (c Song) Module() bot.ModuleInfo {
	return bot.ModuleInfo{
		ID:       "dreamer2q.netease",
		Instance: c,
	}
}

func (Song) Init() {
	logger.Infof("module netease: started")
}

func (Song) Start(bot *bot.Bot) {
	// register event handler
	bot.OnGroupMessage(onGroupMsg)
}

func (Song) Run() {}

func (Song) Stop() {
	logger.Infof("module netease: stopped")
}

var _ bot.Module = Song{}

//onGroupMsg handle specified message and make reply
func onGroupMsg(clt *client.QQClient, msg *message.GroupMessage) {
	text := msg.ToString()
	r := message.NewSendingMessage()

	// rpl.Append(message.NewReply(msg))
	reply := func(fn func() error) {
		if err := fn(); err != nil {
			r.Append(message.NewText(fmt.Sprintf("%v", err)))
		}
		clt.SendGroupMessage(msg.GroupCode, r)
	}

	upload := func(mp3url string) error {
		raw, err := utils.GetBytes(mp3url)
		if err != nil {
			logger.Warnf("upload: %v", err)
			return err
		}
		if !utils.IsAMRorSILK(raw) {
			mt := mimetype.Detect(raw)
			lawful := false
			for _, lt := range lawfulAudioTypes {
				if mt.Is(lt) {
					lawful = true
					break
				}
			}
			if !lawful {
				logger.Infof("invalid audio type: " + mt.String())
				return errors.New("无效类型 " + mt.String())
			}
			silkBytes, err := utils.EncoderSilk(raw)
			if err != nil {
				logger.Warnf("encode: %v", err)
				return err
			}
			raw = silkBytes
		}
		gm, err := clt.UploadGroupPtt(msg.GroupCode, bytes.NewReader(raw))
		if err != nil {
			logger.Warnf("upload group ppt: %v", err)
			return err
		}
		r.Append(gm)
		return nil
	}

	if text == "求歌" {
		reply(func() error {
			song, err := GetRandomSong()
			if err != nil {
				logger.Warnf("求歌: %v", err)
				return err
			}
			return upload(song.Data.Url)
		})
	}

	if strings.HasPrefix(text, "点歌 ") {
		keys := strings.Split(text, " ")
		if len(keys) != 2 {
			return
		}

		reply(func() error {
			song, err := Search(keys[1])
			if err != nil {
				logger.Warnf("search: %v", err)
				return err
			}
			if song.Result.Count == 0 {
				logger.Warnf("search result empty")
				return errors.New("没有找到相应的歌曲")
			}
			songItem := song.Result.Songs[0]
			outerUrl := fmt.Sprintf("http://music.163.com/song/media/outer/url?id=%d", songItem.ID)
			return upload(outerUrl)
		})
	}
}
