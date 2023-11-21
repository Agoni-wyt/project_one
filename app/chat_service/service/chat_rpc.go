package service

import (
	"context"

	// "fmt"

	"one_ser.com/app/chat_service/dao"
	"one_ser.com/app/chat_service/sys_init/model"

	"one_ser.com/app/chat_service/sys_init/redis"
	"one_ser.com/errno"
	"one_ser.com/proto/chat_pb"
)

type ChatSrv struct {
	chat_pb.UnimplementedChatServiceServer
}

func (c *ChatSrv) Send2Client(ctx context.Context, in *chat_pb.SendClientReq) (*chat_pb.SendClientReply, error) {
	//验证是否在线
	// red_key := fmt.Sprintf("%d",in.Data.MessageId)
	//存入redis数据,--msgId
	data := model.Message{
		UserId:   in.UserId,
		TargetId: in.TargetId,
		Type:     int(in.Data.MsgType),
		MsgId:    in.Data.MessageId,
		Content:  in.Data.Content,
	}
	/// redis需要序列化为json存储自定义结构

	redis.MsgAddRedis(ctx, in.TargetId, data)
	res := &chat_pb.SendClientReply{
		Code: errno.SUCCESS,
		Msg:  "success",
		Data: in.Data,
	}
	return res, nil
}

func (c *ChatSrv) CreatCommunity(ctx context.Context, in *chat_pb.CreateCommunityReq) (*chat_pb.CreateCommunityReply, error) {
	resp := new(chat_pb.CreateCommunityReply)
	//验证ownerID是否存在，存在--
	community, err := dao.CommunityCreate(ctx, in)
	if err != nil {
		resp.Code = errno.ERROR
		resp.Msg = err.Error()
		return resp, err
	}
	resp.Code = errno.SUCCESS
	resp.CommunityId = community.CommunityId

	return resp, err
}

func (c *ChatSrv) QuitCommunity(ctx context.Context, in *chat_pb.QuitCommunityReq) (*chat_pb.QuitCommunityReply, error) {
	resp := new(chat_pb.QuitCommunityReply)
	//检查是否为owner
	communityInfo, err := dao.FindCommunityByCommunityId(ctx, in.CommunityId)
	if err != nil {
		resp.Code = errno.ERROR
		resp.Msg = err.Error()
		return resp, err
	}
	if communityInfo.OwnerId == in.UserId {
		//是群主，群主退出将解散该群
		err := dao.DelAllContacts(ctx, in.CommunityId)
		if err != nil {
			resp.Code = errno.ERROR
			resp.Msg = err.Error()
			return resp, err
		}
		//并且删除该群
		err = dao.DelCommunity(ctx, in.UserId, in.CommunityId)
		if err != nil {
			resp.Code = errno.ERROR
			resp.Msg = err.Error()
			return resp, err
		}
		resp.Code = errno.SUCCESS
		resp.Msg = "db delete success"
		return resp, err

	}
	//不是群组
	//删除群组关系
	err = dao.DelContact(ctx, in.UserId, in.CommunityId, 2)
	if err != nil {
		resp.Code = errno.ERROR
		resp.Msg = err.Error()
		return resp, err
	}
	resp.Code = errno.SUCCESS
	resp.Msg = "db delete success"
	return resp, err
}

func (c *ChatSrv) JoinCommunity(ctx context.Context, in *chat_pb.JoinCommunityReq) (*chat_pb.JoinCommunityReply, error) {
	resp := new(chat_pb.JoinCommunityReply)
	//判断community是否存在
	community, err := dao.FindCommunityByCommunityId(ctx, in.CommunityId)
	if err != nil {
		resp.Code = errno.ERROR
		resp.Msg = err.Error()
		return resp, err
	}
	//社区存在
	if community != nil {
		//判断该人物与community关系是否存在
		contact, err := dao.FindContactByUserIdComId(ctx, in.UserId, in.CommunityId)
		if err != nil {
			resp.Code = errno.ERROR
			resp.Msg = err.Error()
			return resp, err
		}
		//关系存在
		if contact.Type == 2 {
			resp.Code = errno.ERROR
			resp.Msg = "u are in community,cant join this community"
		}

	}
	//添加关系
	_, err = dao.AddContact(ctx, in.UserId, in.CommunityId)
	if err != nil {
		resp.Code = errno.ERROR
		resp.Msg = err.Error()
		return resp, err
	}
	resp.Code = errno.SUCCESS
	resp.Msg = "db contact create success"
	return resp, err
}
