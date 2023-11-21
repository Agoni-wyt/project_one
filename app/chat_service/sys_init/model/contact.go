package model

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	UserId int64 //谁的关系
	TargetId int64 //对应的人
	Type int //对应类型 1好友 2群
	Desc string //补充
}
func (Contact)TableName() string {
	return "contact"
}