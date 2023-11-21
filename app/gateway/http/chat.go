package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	// "one_ser.com/app/gateway/rpc"
	// "one_ser.com/app/gateway/rpc/conversion"
	"one_ser.com/app/gateway/http/req"
	"one_ser.com/app/gateway/rpc"
	"one_ser.com/app/gateway/rpc/conversion"
	"one_ser.com/app/gateway/ws"
	"one_ser.com/resp"
	// "one_ser.com/utils"
)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendClientMsg(c *gin.Context) {
	w := c.Writer
	r := c.Request
	conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println(err)
        return
    }
    // defer conn.Close()
	query := r.URL.Query()
	uId := query.Get("userId")
	userId, _ := strconv.ParseInt(uId, 10, 64)
	// tId := query.Get("targetId")
	// targetId,_:=strconv.ParseInt(tId, 10, 64)

	// messageId :=utils.GenID()

	node := ws.NewNode(conn)
	ws.AddMap(userId,node)

	go ws.Send(node) //监听发送
	go ws.Read(node) //收到消息，进行额外操作


	ws.Sys_sendMsg(userId,[]byte("欢迎进入聊天"))

}
func CreatCommunity(c *gin.Context){
	req := req.Create_community{}
	err:=c.ShouldBind(&req)
	fmt.Println(req,err)
	if err != nil {
		resp.ResFail(c.Writer,err.Error())
		return
	}
	// data,err:=service.PlayerRegister(c,req) // --更改为调用rpc文件中的方法
	create_community_pb:=conversion.CreatCommunity_2_CreateCommunityReq(req)

	data,err := rpc.CreateCommunity(c,create_community_pb)

	if err != nil {
		resp.ResFail(c.Writer,err.Error())
		return
	}

	resp.ResOK(c.Writer,data,"创建成功")
	// c.JSON(http.StatusOK, req)

}