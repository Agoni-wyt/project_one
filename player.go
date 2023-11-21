package service

// import (
// 	"context"
// 	"errors"

// 	"one_ser.com/app/gateway/http/req"
// 	"one_ser.com/app/player_service/sys_init/model"
// 	"one_ser.com/app/player_service/dao"
// 	"one_ser.com/app/player_service/service/res"
// 	"one_ser.com/jwt"
// )

//service


// func PlayerRegister (ctx context.Context,req req.RegisterReq)(data *model.Players,err error){
// 	userName:= req.UserName
// 	phone := req.Phone
// 	password := req.Pwd

// 	// 验证是否已经注册，验证密码可行性
// 	if len(password)<=3 {
// 		return nil,errors.New("Password is not long enough")
// 	}

// 	player,error := dao.FindUserByPhone(ctx,phone)
// 	player,error = dao.FindUserByName(ctx,userName)
// 	if player.Phone != "" {
// 		return nil,errors.New("this player is exist")
// 	}
// 	if error!= nil{
// 		return nil,error
// 	}
// 	// 数据库存储
// 	player,err = dao.UserRegister(ctx,req)
// 	if err!= nil{
// 		return nil,err
// 	}
// 	return player,nil

// }

// func PlayerLogin(ctx context.Context,req req.LoginReq)(data *res.TokenData,err error){
// 	player,err := dao.UserLogin(ctx,req)
// 	if err!= nil{
// 		return nil,err
// 	}
// 	player.Password = "" //隐藏密码
// 	token, err := jwt.GenerateToken(player.UserId)
// 	if err!= nil{
// 		return nil,err
// 	}
// 	data = &res.TokenData{
// 		Token: token,
// 		User: player,
// 	}
// 	return data,err
// }

// func GetUserInfo(ctx context.Context,userId int64)(data *model.Players,err error){

// 	player,err := dao.FindUserByUserId(ctx,userId)
// 	if err!= nil{
// 		return nil,err
// 	}
// 	return player,nil

// }