package logging

import (
	"time"
)

type Group struct {
	ID int64 `gorm:"primaryKey"`

	Name            string
	Memo            string
	OwnerID         int64
	GroupCreateTime uint32
	GroupLevel      uint32
	MemberCount     uint16
	MaxMemberCount  uint16
}

type Friend struct {
	ID int64 `gorm:"primaryKey"`

	Remark    string
	Sex       byte
	Age       uint8
	Nickname  string
	Level     int32
	City      string
	Sign      string
	Mobile    string
	LoginDays int64
	Qid       string
}

type MessageBase struct {
	// index
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time

	// common fields
	Time       int32
	MessageId  int32
	InternalId int32
	SenderID   int64
	Message    string
}

type PrivateMessage struct {
	MessageBase
	Self   int64
	Target int64
}

type GroupMessage struct {
	MessageBase
	GroupCode int64
}
