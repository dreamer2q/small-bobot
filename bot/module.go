package bot

import (
	"fmt"
	"sync"
)

var (
	modules   = make(map[string]ModuleInfo)
	modulesMu sync.RWMutex
)

// Module MiraiGo 中的模块
// 用于进行模块化设计
type Module interface {
	Module() ModuleInfo

	// Module 的生命周期

	// Init 初始化, 整个程序中只执行移一次
	// config 会在此之前初始化完成
	Init()

	// Start 应该向Bot注册事件, 不可阻塞
	Start(bot *Bot)

	// Run 通过goroutine启动, 可阻塞
	Run()

	// Stop 应用结束时对所有 Module 进行通知, 在此进行资源回收
	// 这里应该取消在Serve()中订阅的事件，但是MiraiGo暂时没有提供取消订阅的功能
	Stop()
}

// RegisterModule - 向全局添加 Module
func RegisterModule(instance Module) {
	mod := instance.Module()

	if mod.ID == "" {
		panic("module ID missing")
	}
	if mod.Instance == nil {
		panic("missing ModuleInfo.Instance")
	}

	modulesMu.Lock()
	defer modulesMu.Unlock()

	if _, ok := modules[string(mod.ID)]; ok {
		panic(fmt.Sprintf("module already registered: %s", mod.ID))
	}
	modules[string(mod.ID)] = mod
}

// GetModule - 获取一个已注册的 Module 的 ModuleInfo
func GetModule(name string) (ModuleInfo, error) {
	modulesMu.Lock()
	defer modulesMu.Unlock()
	m, ok := modules[name]
	if !ok {
		return ModuleInfo{}, fmt.Errorf("module not registered: %s", name)
	}
	return m, nil
}
