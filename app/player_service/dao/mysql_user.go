package dao

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
	// "one_ser.com/app/gateway/http/req"
	"one_ser.com/app/player_service/sys_init/model"
	"one_ser.com/app/player_service/sys_init/mysql"
	"one_ser.com/proto/player_pb"
	"one_ser.com/utils"
)

var salt string = "secret"

func FindUserByPhone(ctx context.Context, phone string) (*model.Players, error) {
	var userInfo *model.Players
	err := mysql.DB.WithContext(ctx).Model(&model.Players{}).Where("phone = ?", phone).Find(&userInfo).Error
	//可能是数据库错误，也可能是查询不到，当不是查询不到的时候
	if err != nil && err != gorm.ErrEmptySlice {
		return nil, errors.New("query db dailed")
	}
	return userInfo, nil
}
func FindUserByName(ctx context.Context, userName string) (*model.Players, error) {
	var userInfo *model.Players
	err := mysql.DB.WithContext(ctx).Model(&model.Players{}).Where("user_name = ?", userName).Find(&userInfo).Error
	//可能是数据库错误，也可能是查询不到，当不是查询不到的时候
	if err != nil && err != gorm.ErrEmptySlice {
		return nil, errors.New("query db dailed")
	}
	return userInfo, nil
}
func FindUserByUserId(ctx context.Context, userId int64) (*model.Players, error) {
	var userInfo *model.Players
	err := mysql.DB.WithContext(ctx).Model(&model.Players{}).Where("user_id = ?", userId).Find(&userInfo).Error
	//可能是数据库错误，也可能是查询不到，当不是查询不到的时候
	if err != nil && err != gorm.ErrEmptySlice {
		return nil, errors.New("query db dailed")
	}
	return userInfo, nil
}
func UpdateUserByUserId(ctx context.Context,userId int64, filed string, value any) (*model.Players, error) {
	var userInfo *model.Players
	err := mysql.DB.WithContext(ctx).Model(&userInfo).Where("user_id = ?", userId).Update(filed, value).Error
	if err != nil {
		return nil, errors.New("update db dailed")
	}
	userInfo,err = FindUserByUserId(ctx,userId)
	if err != nil {
		return nil, errors.New("query db dailed")
	}
	return userInfo,nil

}

func UserRegister(ctx context.Context, req *player_pb.UserRequest) (*model.Players, error) {
	userId := utils.GenID() //雪花算法生成用户id

	userInfo := &model.Players{
		Phone:        req.Phone,
		UserName:     req.UserName,
		UserId:       userId,
		Password:     utils.MakePassword(req.Password, salt),
		LoginTime:    time.Now(),
		LoginOutTime: time.Now(),
	}
	err := mysql.DB.WithContext(ctx).Create(&userInfo).Error

	if err != nil {
		return nil, errors.New("create db dailed")
	}
	return userInfo, nil
}
func UserLogin(ctx context.Context, req *player_pb.UserRequest) (userInfo *model.Players, err error) {
	if req.UserName != "" {
		userInfo, err = FindUserByName(ctx, req.UserName)
	} else {
		userInfo, err = FindUserByPhone(ctx, req.Phone)
	}
	if err != nil {
		return nil, errors.New("db query dailed")
	}
	if userInfo.Phone == "" {
		return nil, errors.New("user is not exist")
	}
	validation := utils.ValidPassword(req.Password, salt, userInfo.Password)
	if validation {
		userInfo,err = UpdateUserByUserId(ctx,userInfo.UserId,"login_time",time.Now())
		if err !=nil{
			return nil,errors.New("update login_time err")
		}
		return userInfo, nil
	} else {
		return nil, errors.New("wrong password")
	}
}
 func CheckByGoodsIdUserId(ctx context.Context,userId int64,goodsId int64)(*model.Backpack,error){
	var goodsInfo *model.Backpack
	err := mysql.DB.WithContext(ctx).Table("backpack").Where("user_id = ? AND goods_id = ?", userId,goodsId).Find(&goodsId).Error
	if err != nil && err != gorm.ErrEmptySlice {
		return nil, err
	}
	// if err == gorm.ErrEmptySlice {
	// 	return nil, errors.New("this goods is none")
	// }
	return goodsInfo,err
 }
 func UpdateBackpack(ctx context.Context,userId,goodsId int64, filed string, value int64)(*model.Backpack,error){
	var goodsInfo *model.Backpack
	err := mysql.DB.WithContext(ctx).Model(&goodsInfo).Where("user_id = ? AND goods_id = ?", userId,goodsId).Update(filed, value).Error
	if err != nil {
		return nil, err
	}
	goodsInfo,err = CheckByGoodsIdUserId(ctx,goodsInfo.UserId,goodsInfo.GoodsId)
	if err != nil {
		return nil, err
	}
	return goodsInfo,nil
 }
func CreateBackepack(ctx context.Context,req *player_pb.BackpackAddRequest)(*model.Backpack,error){
	goodsInfo:= &model.Backpack{
		UserId: req.UserId,
		GoodsId: req.GoodsId,
		Amount: req.Amount,
		LastPrice: req.LastPrice,
	}
	err := mysql.DB.WithContext(ctx).Create(&goodsInfo).Error
	if err != nil {
		return nil, errors.New("create db dailed")
	}
	return goodsInfo, nil

}
func FindAllGoodsIdByUserId(ctx context.Context,userId int64)([]*model.Backpack,error){
	var allgoods []*model.Backpack
	err := mysql.DB.WithContext(ctx).Where("user_id = ?",userId).Find(&allgoods).Error
	if err != nil {
		return nil, errors.New("query db dailed")
	}
	return allgoods,err
}