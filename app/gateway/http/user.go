package http

import (
	// "net/http"

	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"one_ser.com/app/gateway/http/req"
	"one_ser.com/app/gateway/rpc"
	"one_ser.com/app/gateway/rpc/conversion"

	// "one_ser.com/app/player_service/service"
	"one_ser.com/resp"
)

// 参数校验

func UserRegister(c *gin.Context) {
	req := req.RegisterReq{}
	err:=c.ShouldBind(&req)
	fmt.Println(req,err)
	if err != nil {
		resp.ResFail(c.Writer,err.Error())
		return
	}

	// data,err:=service.PlayerRegister(c,req) // --更改为调用rpc文件中的方法
	regiter_pb:=conversion.RegisterReq_2_UserRequest(req)

	data,err := rpc.UserRegister(c,regiter_pb)

	if err != nil {
		resp.ResFail(c.Writer,err.Error())
		return
	}

	resp.ResOK(c.Writer,data,"注册成功")
	// c.JSON(http.StatusOK, req)

}
// 电话或用户名登录
func Userlogin(c *gin.Context) {
	req := req.LoginReq{}
	err:=c.ShouldBind(&req)
	if err != nil {
		resp.ResFail(c.Writer,err.Error())
		return
	}
	login_pb := conversion.LoginReq_2_UserRequest(req)

	data,err := rpc.UserLogin(c,login_pb)

	if err != nil {
		resp.ResFail(c.Writer,err.Error())
		return
	}
	resp.ResOK(c.Writer,data,"登录成功")

}

func GetUserInfo(c *gin.Context) {
	userIdStr := c.Query("id")
	userId,err := strconv.Atoi(userIdStr)
	if err != nil {
		resp.ResFail(c.Writer,err.Error())
		return
	}
	userId64 := int64(userId)
	userInfo_pb :=conversion.UserInfoReq_2_UserRequest(userId64)
	data,err:= rpc.GetUserInfo(c,userInfo_pb)
	// data,err:= service.GetUserInfo(c,int64(userId))
	if err != nil {
		resp.ResFail(c.Writer,err.Error())
		return
	}
	resp.ResOK(c.Writer,data,"查询信息成功")
}

func BackPackAdd(c *gin.Context) {
	req := req.AddBackPack{}
	err:=c.ShouldBind(&req)
	if err != nil {
		resp.ResFail(c.Writer,err.Error())
		return
	}
	backpackAdd_pb := conversion.BackpackAddReq_2_UserRequest(req)
	data,err := rpc.BackPackAdd(c,backpackAdd_pb)
	if err != nil {
		resp.ResFail(c.Writer,err.Error())
		return
	}
	resp.ResOK(c.Writer,data,"添加物品成功")
}

func BackPackGetAll(c *gin.Context) {
	userIdStr := c.Query("id")
	userId,err := strconv.Atoi(userIdStr)
	if err != nil {
		resp.ResFail(c.Writer,err.Error())
		return
	}
	userId64 := int64(userId)
	userInfo_pb :=conversion.UserInfoReq_2_UserRequest(userId64)
	data,err := rpc.BackPackGetAll(c,userInfo_pb)
	if err != nil {
		resp.ResFail(c.Writer,err.Error())
		return
	}

	resp.ResOK(c.Writer,data.Data,"查找道具")
}