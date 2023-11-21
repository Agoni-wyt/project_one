package rpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"one_ser.com/config"
	"one_ser.com/etcd"
	"one_ser.com/proto/chat_pb"
	"one_ser.com/proto/player_pb"
)

var (
	Register   *etcd.Resolver
	ctx        context.Context
	CancelFunc context.CancelFunc
	UserClient player_pb.UserServiceClient
	ChatClient chat_pb.ChatServiceClient
)

func Init() {
	Register = etcd.NewResolver(config.Conf.Etcd.Endpoints, logrus.New())
	resolver.Register(Register)
	ctx, CancelFunc = context.WithTimeout(context.Background(), 3*time.Second)

	defer Register.Close()
	initClient("player", &UserClient)
	initClient("chat", &ChatClient)
}

func initClient(serviceName string, client interface{}) {
	conn, err := connectServer(serviceName)

	if err != nil {
		panic(err)
	}


	switch c := client.(type) {
	case *player_pb.UserServiceClient:
		*c = player_pb.NewUserServiceClient(conn)
	case *chat_pb.ChatServiceClient:
		*c = chat_pb.NewChatServiceClient(conn)
	default:
		panic("unsupported client type")
	}

}

func connectServer(serviceName string) (conn *grpc.ClientConn, err error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	addr := fmt.Sprintf("%s:/%s", Register.Scheme(), serviceName)

	// Load balance
	if config.Conf.Services[serviceName].LoadBalance {
		log.Printf("load balance enabled for %s\n", serviceName)
		opts = append(opts, grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, "round_robin")))
	}
	conn, err = grpc.DialContext(ctx, addr, opts...)
	return
}
