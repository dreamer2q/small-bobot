# MiraiGo-Template

A template for MiraiGo

[![Go Report Card](https://goreportcard.com/badge/github.com/dreamer2q/MiraiGo-Template)](https://goreportcard.com/report/github.com/dreamer2q/MiraiGo-Template)

基于 [MiraiGo](https://github.com/Mrs4s/MiraiGo) 的多模块组合设计

包装了基础功能,同时设计了一个~~良好~~的项目结构

## 基础配置

账号配置[config.yaml](config/config.yaml)

```yaml
bot:
  # 账号
  account: 1234567
  # 密码
  password: example
```

## 添加一个 Module

添加一个`Module`其实就是实现`Bot`里面定义的`Module`接口

**过程**

1.  定义自己的结构体
2.  实现`Module`接口
3.  向`bot`注册你的`Module`
4.  通过`_`引用来开启你的`Module`

```go
package  YourModule
// 定义一个自己的Module
type YourStruct struct {
	// ...
}

func init() {
    //实例化一个结构体
    instance := &YourStruct{}
    //注册你的Module
	bot.RegisterModule(instance)
}

// 通过这个来让IDE帮你自动实现缺失的方法
var _ bot.Module = YourStruct{}
```

## Module 的生命周期

- `Init()` //Module 初始化，完成 Module 的初始化任务
- `Serve(*bot.Bot)` //Module 注册函数，获取\*bot.Bot 实例， 用于事件注册
- `Start(*bot.Bot)` //Module 主函数，跑在一个新 goroutine 下
- `Stop(bot *Bot, wg *sync.WaitGroup)` //Module 结束函数，对 Module 资源的释放

### 内置 Module

- logging
  将收到的消息按照格式输出至 os.stdout

### 第三方 Module

欢迎 PR

- [logiase.autoreply](https://github.com/Logiase/MiraiGo-module-autoreply)
  按照收到的消息进行回复

# 进阶内容

## Docker 支持

参照 [Dockerfile](./Dockerfile)

# 引入的第三方 go module

- [MiraiGo](https://github.com/Mrs4s/MiraiGo)
  核心协议库
- [viper](https://github.com/spf13/viper)
  用于解析配置文件，同时可监听配置文件的修改
- [logrus](https://github.com/sirupsen/logrus)
  功能丰富的 Logger
- [asciiart](https://github.com/yinghau76/go-ascii-art)
  用于在 console 显示图形验证码
