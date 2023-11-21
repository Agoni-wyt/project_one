package model

import (
	"time"

	"gorm.io/gorm"
)

type Players struct {
	gorm.Model //嵌入默认字段

	Phone     string    `gorm:"not null;unique"`	//用户名
	Password  string	`gorm:"not null"`	//密码

	UserId       int64  `gorm:"not null;unique;"`   //用户id
	UserName     string  `gorm:"not null;default:''"`  //用户名
	Sex          int8    `gorm:"not null;default:0"`  //0未知 1女 2男
	Status       int8    `gorm:"not null;default:0"`  //0离线 1在线
	Avatar       string  `gorm:"not null;default:''"`  //头像
	Assets1Num   int64	 `gorm:"not null;default:100"`  //资产1数量
	Assets2Num   int64	 `gorm:"not null;default:0"`  //资产2数量
	TradeNum	 int64   `gorm:"not null;default:0"`  //交易次数
	TitleId		 int8	 `gorm:"not null;default:0"`  //称号Id
	LoginTime	 time.Time //最近登入时间
	LoginOutTime time.Time // 上次离线时间
}

func (Players) TableName() string {
	return "players"
}