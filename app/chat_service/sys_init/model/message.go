package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model       //嵌入默认字段
	MsgId      int64 `gorm:"not null;unique;"`//消息Id
	UserId     int64 //发送者用户id
	TargetId   int64 //接收者id
	Type       int    //发送类型  1私聊  2群聊  3心跳
	Media      int    //消息类型  1文字 2表情包 3语音 4图片 /表情包
	Content    string //消息内容
	CreateTime uint64 //创建时间
	ReadTime   uint64 //读取时间
}

func (Message) TableName() string {
	return "message"
}
