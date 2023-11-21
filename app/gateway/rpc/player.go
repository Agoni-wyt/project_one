package rpc

import (
	"context"
	"errors"

	"one_ser.com/errno"
	"one_ser.com/proto/player_pb"
)

// var (
// 	UserClient player_pb.UserServiceClient
// )

func UserRegister(ctx context.Context, req *player_pb.UserRequest) (resp *player_pb.UserDetailResponse, err error) {
	resp, err = UserClient.UserRegister(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.Code != errno.SUCCESS {
		err = errors.New(errno.GetMsg(int(resp.Code)))
		return nil, err
	}

	return resp, nil
}

func UserLogin(ctx context.Context,req *player_pb.UserRequest)(resp *player_pb.UserDetailResponse, err error) {
	resp, err = UserClient.UserLogin(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.Code != errno.SUCCESS {
		err = errors.New(errno.GetMsg(int(resp.Code)))
		return nil, err
	}

	return resp, nil
}

func GetUserInfo(ctx context.Context, req *player_pb.UserInfo)(resp *player_pb.UserDetailResponse, err error) {
	resp, err = UserClient.UserInfoGet(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.Code != errno.SUCCESS {
		err = errors.New(errno.GetMsg(int(resp.Code)))
		return nil, err
	}

	return resp, nil
}

func BackPackAdd(ctx context.Context,req *player_pb.BackpackAddRequest)(resp *player_pb.BackpackgoodsResponse,err error){
	resp, err = UserClient.BackpackAdd(ctx,req)
	if err != nil {
		return nil, err
	}

	if resp.Code != errno.SUCCESS {
		err = errors.New(errno.GetMsg(int(resp.Code)))
		return nil, err
	}

	return resp, nil
}

func BackPackGetAll(ctx context.Context,req *player_pb.UserInfo)(resp *player_pb.BackpackGetAllResponse,err error){
	resp, err = UserClient.BackpackGetAll(ctx,req)
	if err != nil {
		return nil, err
	}

	if resp.Code != errno.SUCCESS {
		err = errors.New(errno.GetMsg(int(resp.Code)))
		return nil, err
	}
	return resp, nil
}