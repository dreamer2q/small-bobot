package bot

import "strings"

// ModuleID 模块ID
// 请使用 小写 并用 _ 代替空格
// Example:
// - logiase.autoreply
type ModuleID string

// Namespace - 获取一个 Module 的 Namespace
func (id ModuleID) Namespace() string {
	lastDot := strings.LastIndex(string(id), ".")
	if lastDot < 0 {
		return ""
	}
	return string(id)[:lastDot]
}

// Name - 获取一个 Module 的 Name
func (id ModuleID) Name() string {
	if id == "" {
		return ""
	}
	parts := strings.Split(string(id), ".")
	return parts[len(parts)-1]
}

// ModuleInfo 模块信息
type ModuleInfo struct {
	// ID 模块的名称
	// 应全局唯一
	ID ModuleID

	// Instance 返回 Module
	Instance Module
}

func (mi ModuleInfo) String() string {
	return string(mi.ID)
}
