package welcome

type Welcome struct {
	ID uint `gorm:"primarykey"`

	GroupID int64 // 群号

	Welcome string // 欢迎信息

	Enabled bool // 是否启用
}
