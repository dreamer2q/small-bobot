package faq

import (
	"time"
)

/*
FAQ 问答规则
*/
type Faq struct {
	ID uint `gorm:"primarykey"`

	GroupID  int64  // 群号
	Question string // 正则问题
	Answer   string // 回答

	Enabled bool // 是否启用
}

/*
FAQ 问答日志
*/
type FaqLog struct {
	FaqID     uint      `gorm:"index"` // 关联的 FAQ 规则
	CreatedAt time.Time // 记录日期

	UID int64  // 提问者 QQ
	Msg string // 提问内容
}
