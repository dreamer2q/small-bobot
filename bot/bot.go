package bot

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"os"
	"strings"
	"sync"

	"github.com/Mrs4s/MiraiGo/client"
	"github.com/sirupsen/logrus"
	asc2art "github.com/yinghau76/go-ascii-art"
)

// Bot 全局 Bot
type Bot struct {
	*client.QQClient

	start bool
}

// Instance Bot 实例
var Instance *Bot

var logger = logrus.WithField("bot", "internal")

type Config struct {
	Account  int64
	Password string
	Device   []byte
}

// Init 快速初始化
func Init(conf Config) {
	Instance = &Bot{
		client.NewClient(
			conf.Account,
			conf.Password,
		),
		false,
	}
	if conf.Device == nil {
		logger.Debugf("no device specified, generate random device")
		client.GenRandomDevice()
	} else {
		err := client.SystemDeviceInfo.ReadJson(conf.Device)
		if err != nil {
			logger.WithError(err).Panic("parse device error")
		}
	}
}

// Login 登录
func Login() {
	resp, err := Instance.Login()
	console := bufio.NewReader(os.Stdin)

	for {
		if err != nil {
			logger.WithError(err).Fatal("unable to login")
		}

		var text string
		if !resp.Success {
			switch resp.Error {
			case client.SliderNeededError:
				if client.SystemDeviceInfo.Protocol == client.AndroidPhone {
					logger.Warn("Android Phone Protocol DO NOT SUPPORT Slide verify")
					logger.Warn("please use other protocol")
					os.Exit(2)
				}
				Instance.AllowSlider = false
				Instance.Disconnect()
				resp, err = Instance.Login()
				continue
			case client.NeedCaptcha:
				img, _, _ := image.Decode(bytes.NewReader(resp.CaptchaImage))
				fmt.Println(asc2art.New("image", img).Art)
				fmt.Print("please input captcha: ")
				text, _ := console.ReadString('\n')
				resp, err = Instance.SubmitCaptcha(strings.ReplaceAll(text, "\n", ""), resp.CaptchaSign)
				continue
			case client.SMSNeededError:
				fmt.Println("device lock enabled, Need SMS Code")
				fmt.Printf("Send SMS to %s ? (yes)", resp.SMSPhone)
				t, _ := console.ReadString('\n')
				t = strings.TrimSpace(t)
				if t != "yes" {
					os.Exit(2)
				}
				if !Instance.RequestSMS() {
					logger.Warnf("unable to request SMS Code")
					os.Exit(2)
				}
				logger.Warn("please input SMS Code: ")
				text, _ = console.ReadString('\n')
				resp, err = Instance.SubmitSMS(strings.ReplaceAll(strings.ReplaceAll(text, "\n", ""), "\r", ""))
				continue
			case client.SMSOrVerifyNeededError:
				fmt.Println("device lock enabled, choose way to verify:")
				fmt.Println("1. Send SMS Code to ", resp.SMSPhone)
				fmt.Println("2. Scan QR Code")
				fmt.Print("input (1,2):")
				text, _ = console.ReadString('\n')
				text = strings.TrimSpace(text)
				switch text {
				case "1":
					if !Instance.RequestSMS() {
						logger.Warnf("unable to request SMS Code")
						os.Exit(2)
					}
					logger.Warn("please input SMS Code: ")
					text, _ = console.ReadString('\n')
					resp, err = Instance.SubmitSMS(strings.ReplaceAll(strings.ReplaceAll(text, "\n", ""), "\r", ""))
					continue
				case "2":
					fmt.Printf("device lock -> %v\n", resp.VerifyUrl)
					os.Exit(2)
				default:
					fmt.Println("invalid input")
					os.Exit(2)
				}
			case client.UnsafeDeviceError:
				fmt.Printf("device lock -> %v\n", resp.VerifyUrl)
				os.Exit(2)
			case client.OtherLoginError, client.UnknownLoginError:
				logger.Fatalf("login failed: %v", resp.ErrorMessage)
				os.Exit(3)
			}

		}

		break
	}

	logger.Infof("bot login: %s", Instance.Nickname)
}

// RefreshList 刷新联系人
func RefreshList() {
	logger.Info("start reload friends list")
	err := Instance.ReloadFriendList()
	if err != nil {
		logger.WithError(err).Error("unable to load friends list")
	}
	logger.Infof("load %d friends", len(Instance.FriendList))
	logger.Info("start reload groups list")
	err = Instance.ReloadGroupList()
	if err != nil {
		logger.WithError(err).Error("unable to load groups list")
	}
	logger.Infof("load %d groups", len(Instance.GroupList))
}

// StartService 启动服务
// 根据 Module 生命周期 此过程应在Login前调用
// 请勿重复调用
func StartService() {
	if Instance.start {
		return
	}

	Instance.start = true

	logger.Infof("initializing modules ...")
	for _, mi := range modules {
		mi.Instance.Init()
	}

	logger.Info("all modules initialized")

	logger.Info("registering modules serve functions ...")
	for _, mi := range modules {
		mi.Instance.Serve(Instance)
	}
	logger.Info("all modules serve functions registered")

	logger.Info("starting modules tasks ...")
	for _, mi := range modules {
		go mi.Instance.Start(Instance)
	}
	logger.Info("tasks running")
}

// Stop 停止所有服务
// 调用此函数并不会使Bot离线
func Stop() {
	logger.Warn("stopping ...")
	wg := sync.WaitGroup{}
	for _, mi := range modules {
		wg.Add(1)
		mi.Instance.Stop(Instance, &wg)
	}
	wg.Wait()
	logger.Info("stopped")
	modules = make(map[string]ModuleInfo)
}
