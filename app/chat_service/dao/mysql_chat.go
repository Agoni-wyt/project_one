package dao

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
	"one_ser.com/app/chat_service/sys_init/model"
	"one_ser.com/app/chat_service/sys_init/mysql"
	"one_ser.com/proto/chat_pb"
)

func CommunityCreate(ctx context.Context,req *chat_pb.CreateCommunityReq)(*model.Community,error){
	data := &model.Community{
		CommunityId: req.CommunityId,
		OwnerId: req.OwnerId,
		Name: req.Name,
		LimitNum: req.LimitNum,
		Desc: req.Desc,
		CreateTime: time.Now(),
	}
	err:=mysql.DB.WithContext(ctx).Create(data).Error
	if err != nil {
		return nil, errors.New("create db dailed")
	}
	return data, nil
}

func FindCommunityByCommunityId(ctx context.Context,communityId int64)(*model.Community,error){
	var communityInfo *model.Community
	err:= mysql.DB.WithContext(ctx).Model(&model.Community{}).Where("community_id = ?",communityId).Find(&communityInfo).Error
	if err != nil && err != gorm.ErrEmptySlice {
		return nil, errors.New("query db dailed")
	}
	return communityInfo, nil
}

func FindCommunitiesByOwenId(ctx context.Context,ownerId int64)([]*model.Community,error){
	var communities []*model.Community
	err:=mysql.DB.WithContext(ctx).Where("owner_id = ?",ownerId).Find(&communities).Error
	if err != nil {
		return nil, errors.New("query db dailed")
	}
	return communities,err

}
func DelCommunity(ctx context.Context,userId,communityId int64)error{
	err:= mysql.DB.WithContext(ctx).Where(&model.Community{OwnerId: userId,CommunityId: communityId}).Delete(&model.Community{})
	if err != nil {
		return errors.New("delete owner's community err ")
	}
	return nil
}
func FindContactByUserIdComId(ctx context.Context,userId,communityId int64)(*model.Contact,error){
	var contact *model.Contact
	err:= mysql.DB.WithContext(ctx).Where(&model.Contact{UserId: userId,TargetId: communityId,Type: 2}).Find(&contact).Error
	if err != nil {
		return nil, err
	}
	return contact,nil
}
func AddContact(ctx context.Context,userId,communityId int64)(*model.Contact,error){
	data := &model.Contact{
		UserId: userId,
		TargetId: communityId,
		Type: 2,
		Desc: "群关系",
	}
	err:=mysql.DB.WithContext(ctx).Create(data).Error
	if err != nil {
		return nil, errors.New("create db dailed")
	}
	return data, nil
}


func DelContact(ctx context.Context,userId,targetId int64,contactType int)error{
	err:= mysql.DB.WithContext(ctx).Where(&model.Contact{UserId: userId,TargetId: targetId,Type: contactType}).Delete(&model.Contact{}).Error
	if err != nil {
		return errors.New(fmt.Sprintf("delete db dailed is %f", err))
	}
	return nil
}
func DelAllContacts(ctx context.Context,communityId int64)error{
	var contacts[]*model.Contact
	err:= mysql.DB.WithContext(ctx).Where(&model.Contact{TargetId: communityId,Type: 2}).Find(&contacts).Error
	if err != nil {
		return errors.New("find contact err")
	}
	err = mysql.DB.WithContext(ctx).Delete(&contacts).Error
	if err != nil {
		return errors.New("del all community contacts err")
	}
	return nil
}

func AllMemberInCommunity(ctx context.Context,communityId int64)([]int64,error){
	var contacts[]*model.Contact
	var userIds []int64
	err:= mysql.DB.WithContext(ctx).Where(&model.Contact{TargetId: communityId,Type: 2}).Find(&contacts).Error
	if err != nil {
		return nil, errors.New("find contact err")
	}
	for _, v := range contacts {
		userIds = append(userIds, v.UserId)
	}
	return userIds,nil
}