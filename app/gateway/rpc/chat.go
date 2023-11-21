package rpc

import (
	"context"
	"errors"

	"one_ser.com/errno"
	"one_ser.com/proto/chat_pb"
)


func Send2Client(ctx context.Context, req *chat_pb.SendClientReq)(resp *chat_pb.SendClientReply, err error){
	resp, err = ChatClient.Send2Client(ctx,req)
	if err != nil {
		return nil, err
	}

	if resp.Code != errno.SUCCESS {
		err = errors.New(errno.GetMsg(int(resp.Code)))
		return nil, err
	}

	return resp, nil
}


func CreateCommunity(ctx context.Context,req *chat_pb.CreateCommunityReq)(resp *chat_pb.CreateCommunityReply,err error){
	resp, err = ChatClient.CreateCommunity(ctx,req)
	if err != nil {
		return nil, err
	}

	if resp.Code != errno.SUCCESS {
		err = errors.New(errno.GetMsg(int(resp.Code)))
		return nil, err
	}

	return resp, nil
}


