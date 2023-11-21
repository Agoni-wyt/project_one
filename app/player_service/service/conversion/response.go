package conversion

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	"one_ser.com/app/player_service/sys_init/model"
	"one_ser.com/proto/player_pb"
)

type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

func PlayerModel_2_pd(model *model.Players)(pb *player_pb.UserResponse){
	loginTime := timestamppb.New(model.LoginTime)
	loginOutTime := timestamppb.New(model.LoginOutTime)
	pb = &player_pb.UserResponse{
		UserId: model.UserId,
		UserName: model.UserName,
		Sex: int32(model.Sex),
		Avater: model.Avatar,
		Assets1Num: model.Assets1Num,
		Assets2Num: model.Assets2Num,
		TradeNum: model.TradeNum,
		TitleId: int32(model.TitleId),
		LoginTime: loginTime,
		LoginOutTime: loginOutTime,
	}
	return pb
}