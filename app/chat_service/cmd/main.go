package main

import (
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"one_ser.com/app/chat_service/service"
	"one_ser.com/app/chat_service/sys_init/mysql"
	"one_ser.com/app/chat_service/sys_init/redis"
	"one_ser.com/config"
	"one_ser.com/etcd"
	"one_ser.com/logger"
	"one_ser.com/proto/chat_pb"
	"one_ser.com/utils"
)

// 初始化数据库，以及etcd
func main() {
	//1. 加载配置文件
	var cfn string
	err := config.Init(cfn)
	if err != nil {
		panic(err) //启动时加载配置文件失败，直接退出
	}
	//2. 加载日志
	err = logger.Init(config.Conf.LogConfig, config.Conf.Mode)
	if err != nil {
		panic(err) //启动时加载日志模块失败，直接退出
	}
	// 3. 初始化MySQL
	err = mysql.Init(config.Conf.MySQLConfig)
	if err != nil {
		panic(err) //启动时数据库模块失败，直接退出
	}

	// 4. 初始化redis
	err = redis.Init(config.Conf.RedisConfig)
	if err != nil {
		panic(err) //启动时redis模块失败，直接退出ss
	}

	// 5. 初始化snowflake，初始
	err = utils.SFInit(config.Conf.StartTime, config.Conf.MachineId)
	if err != nil {
		panic(err)
	}

	// etcd 地址
	etcdAddress := config.Conf.Etcd.Endpoints
	// 服务注册
	etcdRegister := etcd.NewRegister(etcdAddress, logrus.New())
	addr:=fmt.Sprintf("%s:%d",config.Conf.IP,config.Conf.Services["chat"].RpcPort)
	grpcAddress := addr
	defer etcdRegister.Stop()
	userNode := etcd.Server{
		Name: "chat",
		Addr: grpcAddress,
	}
	server := grpc.NewServer()
	defer server.Stop()


	// 绑定service
	chat_pb.RegisterChatServiceServer(server, &service.ChatSrv{})
	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		panic(err)
	}
	if _, err := etcdRegister.Register(userNode, 10); err != nil {
		panic(fmt.Sprintf("start server failed, err: %v", err))
	}
	logrus.Info("server started listen on ", grpcAddress)
	if err := server.Serve(lis); err != nil {
		panic(err)
	}

}
