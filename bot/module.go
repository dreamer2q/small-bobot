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

	// Init 初始化
	Init()

	// Serve 向Bot注册服务函数
	Serve(bot *Bot)

	// Start 启用Module
	Start(bot *Bot)

	// Stop 应用结束时对所有 Module 进行通知, 在此进行资源回收
	Stop(bot *Bot, wg *sync.WaitGroup)
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
	//if val := mod.Instance; val == nil {
	//	panic("ModuleInfo.Instance must return a non-nil module instance")
	//}

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
