package service

import (
	"context"
	"errors"

	"one_ser.com/app/player_service/dao"
	"one_ser.com/app/player_service/service/conversion"
	"one_ser.com/errno"
	"one_ser.com/proto/player_pb"
)

type UserSrv struct {
	player_pb.UnimplementedUserServiceServer
}

func (u *UserSrv) UserRegister(ctx context.Context, req *player_pb.UserRequest) (resp *player_pb.UserDetailResponse, err error) {
	resp = new(player_pb.UserDetailResponse)
	userName:= req.UserName
	phone := req.Phone
	password := req.Password

	// 验证是否已经注册，验证密码可行性
	if len(password)<=3 {
		return nil,errors.New("Password is not long enough")
	}

	player,error := dao.FindUserByPhone(ctx,phone)
	player,error = dao.FindUserByName(ctx,userName)
	if player.Phone != "" {
		resp.Code = errno.ERROR
		return resp,errors.New("this player is exist")
	}
	if error!= nil{
		resp.Code = errno.ERROR
		return resp,error
	}
	// 数据库存储
	player,err = dao.UserRegister(ctx,req)
	if err!= nil{
		resp.Code = errno.ERROR
		return resp,err
	}
	resp.Code = errno.SUCCESS
	resp.UserDetail = conversion.PlayerModel_2_pd(player)
	return
}


func (u *UserSrv) UserLogin(ctx context.Context, req *player_pb.UserRequest) (resp *player_pb.UserDetailResponse, err error) {
	resp = new(player_pb.UserDetailResponse)
	player,err := dao.UserLogin(ctx,req)
	if err!= nil{
		resp.Code = errno.ERROR
		return
	}
	player.Password = "" //隐藏密码
	if err!= nil{
		resp.Code = errno.ERROR
		return
	}
	resp.Code = errno.SUCCESS
	resp.UserDetail = conversion.PlayerModel_2_pd(player)

	return


}

func (u *UserSrv) UserInfoGet(ctx context.Context, req *player_pb.UserInfo) (resp *player_pb.UserDetailResponse, err error) {
	resp = new(player_pb.UserDetailResponse)
	userId := req.UserId
	player,err := dao.FindUserByUserId(ctx,userId)
	if err!= nil{
		resp.Code = errno.ERROR
		return resp,err
	}
	resp.Code = errno.SUCCESS
	resp.UserDetail = conversion.PlayerModel_2_pd(player)

	return

}

func (u *UserSrv)BackpackAdd(ctx context.Context, req *player_pb.BackpackAddRequest)(resp *player_pb.BackpackgoodsResponse, err error){
	resp = new(player_pb.BackpackgoodsResponse)
	userId := req.UserId
	player,err := dao.FindUserByUserId(ctx,userId)
	if err!= nil || player==nil{
		resp.Code = errno.ERROR
		return nil,err
	}
	goodsId := req.GoodsId
	amount := req.Amount
	// lastprice := req.LastPrice

	goodsInfo,err := dao.CheckByGoodsIdUserId(ctx,userId,goodsId)
	if err!= nil{
		resp.Code = errno.ERROR
		return
	}
	if goodsInfo!=nil {
		//更新
		goodsInfo,err = dao.UpdateBackpack(ctx,goodsInfo.UserId,goodsInfo.GoodsId,"amount",goodsInfo.Amount+amount)
		if err!= nil{
			resp.Code = errno.ERROR
			return
		}
	}else{
		//创建
		goodsInfo,err = dao.CreateBackepack(ctx,req)
		if err!= nil{
			resp.Code = errno.ERROR
			return
		}
	}
	resp.Code = errno.SUCCESS
	resp.UserGoods = &player_pb.BackpackUserGoodsInfo{
		UserId: goodsInfo.UserId,
		GoodsId: goodsInfo.GoodsId,
		Amount: goodsInfo.Amount,
	}
	return
}

func (u *UserSrv)BackpackGetAll(ctx context.Context, req *player_pb.UserInfo)(resp *player_pb.BackpackGetAllResponse, err error){
	//获取全部--判断用户存在--查询数据是否存在
	resp = new(player_pb.BackpackGetAllResponse)
	userId := req.UserId
	player,err := dao.FindUserByUserId(ctx,userId)
	if err!= nil || player==nil{
		resp.Code = errno.ERROR
		return nil,err
	}
	modelGoodsAll,err := dao.FindAllGoodsIdByUserId(ctx,userId)
	if err!= nil {
		resp.Code = errno.ERROR
		return
	}
	var goodsAll []*player_pb.BackpackGoodsInfo
	resp.Code = errno.SUCCESS
	for _, goods := range modelGoodsAll {
		goodsAll = append(goodsAll, &player_pb.BackpackGoodsInfo{
			GoodsId: goods.GoodsId,
			Amount: goods.Amount,
			LastPrice: goods.LastPrice,
		})
	}
	resp.Data = goodsAll
	return
}