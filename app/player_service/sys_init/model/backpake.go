package model

import "gorm.io/gorm"

type Backpack struct {
	gorm.Model        //嵌入默认字段
	UserId     int64  //用户id
	GoodsId    int64  `gorm:"not null;default:0"`//物品Id
	Amount     int64  `gorm:"not null;default:0"`//数量
	LastPrice  int64  `gorm:"not null;default:0"`//上次交易价格
}

func (Backpack) TableName() string {
	return "backpack"
}
