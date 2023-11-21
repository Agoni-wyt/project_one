package ws

import (
	"context"
	"encoding/json"
	"net"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"gopkg.in/fatih/set.v0"
	"one_ser.com/app/gateway/http/req"
	"one_ser.com/app/gateway/rpc"
	"one_ser.com/app/gateway/rpc/conversion"
	"one_ser.com/config"
	"one_ser.com/proto/chat_pb"
	// "one_ser.com/utils"
)

type Node struct {
	Conn          *websocket.Conn //连接
	Addr          string          //客户端地址
	FirstTime     uint64          //首次连接时间
	HeartbeatTime uint64          //心跳时间
	LoginTime     uint64          //登录时间
	DataQueue     chan []byte     //消息
	GroupSets     set.Interface   //好友 / 群  -->集合
}

// 映射关系
var clientMap map[int64]*Node = make(map[int64]*Node, 0)
var onlineUsers map[int64]bool = make(map[int64]bool, 0)

// 读写锁
var (
	rwLocker sync.RWMutex
	mu       sync.Mutex
)

// 用于获取连接
func NewNode(conn *websocket.Conn) *Node {
	currentTime := uint64(time.Now().Unix())
	return &Node{
		Conn:          conn,
		Addr:          conn.RemoteAddr().String(), //客户端地址
		HeartbeatTime: currentTime,                //心跳时间
		LoginTime:     currentTime,                //登录时间
		DataQueue:     make(chan []byte, 50),
		GroupSets:     set.New(set.ThreadSafe),
	}
}

// userId绑定
func AddMap(userId int64, node *Node) {
	rwLocker.Lock()
	clientMap[userId] = node
	onlineUsers[userId] = true //上线
	rwLocker.Unlock()
}

// 读客户端消息
func Read(node *Node) {
	for {
		messageType, data, err := node.Conn.ReadMessage()
		if err != nil {
			if messageType == -1 && websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure, websocket.CloseNoStatusReceived) {
				panic("websocket is close")

			} else if messageType != websocket.PingMessage {
				return
			}
		}
		msg := req.Chat_msg{}
		err = json.Unmarshal(data, &msg)
		if err != nil {
			log.Error(err)
		}

		if msg.Type == 3 {
			currentTime := uint64(time.Now().Unix())
			node.Heartbeat(currentTime)
		} else {
			dispatch(data) //调度 -> 选择rpc是那种方式发送
			broadMsg(data) //将消息广播到局域网
		}
	}
}

var udpsendChan chan []byte = make(chan []byte, 1024)

// 通过udp协议传输
func broadMsg(data []byte) {
	udpsendChan <- data
}
func init() {
	go udpSendProc()
	go udpRecvProc()
	log.Info("init goroutine ")
}

// 完成udp数据发送协程
func udpSendProc() {
	con, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(10, 100, 139, 246),
		Port: int(config.Conf.Udp),
	})
	defer con.Close()
	if err != nil {
		log.Error(err)
	}
	for {
		select {
		case data := <-udpsendChan:
			_, err := con.Write(data)
			if err != nil {
				log.Error(err)
				return
			}
		}
	}

}

// 完成udp数据接收协程
func udpRecvProc() {
	con, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: int(config.Conf.Udp),
	})
	if err != nil {
		log.Error(err)
	}
	defer con.Close()
	for {
		var buf [512]byte
		n, err := con.Read(buf[0:])
		if err != nil {
			log.Error(err)
			return
		}
		dispatch(buf[0:n])
	}
}

// 后端调度逻辑处理
func dispatch(data []byte) {
	msg := req.Chat_msg{}
	msg.CreateTime = uint64(time.Now().Unix())
	err := json.Unmarshal(data, &msg)
	if err != nil {
		log.Error(err)
		return
	}

	switch msg.Type {
	case 1: //私信
		//grpc->放入redis
		sendClient_pb := conversion.Ws_2_SendClientReq(msg)
		_, err := rpc.Send2Client(context.Background(), sendClient_pb)
		if err != nil {
			log.Error(err)
			return
		}
		Sys_sendMsg(msg.TargetId, data) //rpc --> rpc.chat()
	case 2: //在线广播
		OnlineBroadcast(data) //发送的群ID ，消息内容
		// 	//case 4:
		// 	//
	}

}

func OnlineBroadcast(msg []byte) {
	for i, v := range onlineUsers {
		if v {
			Sys_sendMsg(i, msg)
		}
	}
}

func MarshalBinary(resp *chat_pb.SendClientReply) ([]byte, error) {
	return json.Marshal(&resp)
}
func UnMarshal(data string, model any) {
	err := json.Unmarshal([]byte(data), &model)
	if err != nil {
		log.Error(err)
	}
}

// 更新用户心跳
func (node *Node) Heartbeat(currentTime uint64) {
	node.HeartbeatTime = currentTime
	return
}

func Send(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			if node.Conn != nil {
				err := node.Conn.WriteMessage(websocket.TextMessage, data)
				if err != nil {
					log.Errorf("senderr:%s", err)
					return
				}
			}

		}
	}
}

func Sys_sendMsg(userId int64, msg []byte) {
	rwLocker.RLock()
	node, ok := clientMap[userId] //拿出连接
	rwLocker.RUnlock()

	if ok {
		log.Debug("sendMsg-->User ,userId: %d, msg: %s", userId, string(msg))
		node.DataQueue <- msg //将消息放入队列
	}
}

// 清理超时连接
func CleanConnection(param interface{}) (result bool) {
	result = true
	defer func() {
		if r := recover(); r != nil {
			log.Error("cleanConnection err", r)
		}
	}()
	currentTime := uint64(time.Now().Unix())
	for i := range clientMap {
		node := clientMap[i]
		if node.IsHeartbeatTimeOut(currentTime) {
			log.Error("心跳超时..... 关闭连接：", node)
			node.Conn.Close()
			onlineUsers[i] = false
		}
	}
	return result
}

// 用户心跳是否超时
func (node *Node) IsHeartbeatTimeOut(currentTime uint64) (timeout bool) {
	if node.HeartbeatTime+3 <= currentTime {
		log.Error("心跳超时。。。自动下线", node)
		timeout = true
	}
	return
}
