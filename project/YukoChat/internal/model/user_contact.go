package model

import (
	"time"

	"gorm.io/gorm"
)

// todo: 将常量命名改为纯大写
const (
	ContactTypeUser       = 0 // 联系类型：用户
	ContactTypeGroup      = 1 // 联系类型：群聊
	StatusNormal          = 0 // 联系状态：正常
	StatusBlackList       = 1 // 联系状态：拉黑
	StatusBeBlackList     = 2 // 联系状态：被拉黑
	StatusDeletedFriend   = 3 // 联系状态：删除好友
	StatusBeDeletedFriend = 4 // 联系状态：被删除好友
	StatusMuted           = 5 // 联系状态：被禁言
	StatusQuitGroup       = 6 // 联系状态：退出群聊
	StatusKickedGroup     = 7 // 联系状态：被踢出群聊
)

// 每一对关系都会对应数据库的两行
// A 加 B 好友，会产生 (A, B, ...) 和 (B, A, ...) 两行数据
// A 加 群聊 G, 会产生 (A, G, ...) 一行数据，群聊 G 不会有对应的 (G, A, ...) 数据
type UserContact struct {
	Id          int64          `gorm:"column:id;primaryKey;comment:自增id"`
	UserId      string         `gorm:"column:user_id;index;type:char(20);not null;comment:用户唯一id"`
	ContactId   string         `gorm:"column:contact_id;index;type:char(20);not null;comment:对应联系id"`
	ContactType int8           `gorm:"column:contact_type;not null;comment:联系类型，0.用户，1.群聊"`
	Status      int8           `gorm:"column:status;not null;comment:联系状态，0.正常，1.拉黑，2.被拉黑，3.删除好友，4.被删除好友，5.被禁言，6.退出群聊，7.被踢出群聊"`
	CreatedAt   time.Time      `gorm:"column:created_at;type:datetime;not null;comment:创建时间"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;index;comment:删除时间"`
}
