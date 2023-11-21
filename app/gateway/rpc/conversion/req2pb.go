package conversion

//用于转变变量类型 req -> pb
import (
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"
	"one_ser.com/app/gateway/http/req"
	"one_ser.com/proto/chat_pb"
	"one_ser.com/proto/player_pb"
	"one_ser.com/utils"
)

func RegisterReq_2_UserRequest(req req.RegisterReq) (*player_pb.UserRequest){
	pb := &player_pb.UserRequest{
		Password: req.Pwd,
		Phone: req.Phone,
		UserName: req.UserName,
	}
	return pb
}

func LoginReq_2_UserRequest(req req.LoginReq)(*player_pb.UserRequest){
	pb := &player_pb.UserRequest{
		Password: req.Pwd,
		Phone: req.Phone,
		UserName: req.UserName,
	}

	return pb
}
func UserInfoReq_2_UserRequest(userId int64)(*player_pb.UserInfo){
	pb := &player_pb.UserInfo{
		UserId: userId,
	}
	return pb
}
func BackpackAddReq_2_UserRequest(req req.AddBackPack)(*player_pb.BackpackAddRequest){
	pb := &player_pb.BackpackAddRequest{
		UserId: req.UserId,
		GoodsId: req.GoodsId,
		Amount: req.Amount,
		LastPrice: req.LastPrice,
	}
	return pb
}

func Ws_2_SendClientReq(msg req.Chat_msg)(*chat_pb.SendClientReq){
	messageId := utils.GenID()
	pb:= &chat_pb.SendClientReq{
		UserId: msg.UserId,
		SystemId: fmt.Sprintf("%d", msg.UserId),
		TargetId: msg.TargetId,
		Data: &chat_pb.MessageInfo{
			MessageId: messageId,
			MsgType: int32(msg.Type),
			Content: msg.Content,
			SendTime: timestamppb.Now(),
		},
	}
	return pb
}

func CreatCommunity_2_CreateCommunityReq(msg req.Create_community)(*chat_pb.CreateCommunityReq){
	communityId := utils.GenID()
	pb:= &chat_pb.CreateCommunityReq{
		CommunityId: communityId,
		OwnerId: msg.OwnerId,
		LimitNum: msg.LimitNum,
		Desc:msg.Desc,
	}
	return pb
}