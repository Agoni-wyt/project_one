package model

import (
	"time"

	"gorm.io/gorm"
)

type Community struct {
	gorm.Model         //嵌入默认字段
	CommunityId int64  `gorm:"not null;unique;"`//群组Id
	Name        string `gorm:"not null;default:''"`//群名
	OwnerId     int64  //创建者
	LimitNum    int64  `gorm:"not null;default:100"`//限制人数
	Desc        string `gorm:"not null;default:''"`//解释
	CreateTime  time.Time //创建时间
}

func (Community) TableName() string {
	return "community"
}
