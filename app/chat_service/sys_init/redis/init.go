package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/disgoorg/log"
	"github.com/go-redis/redis/v8"
	"one_ser.com/config"
)

var (
	Red *redis.Client
)

func Init(cfg *config.RedisConfig) (err error) {
	ctx := context.Background()
	Red = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})
	_, err = Red.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}
//Publish 发布消息到Redis
func Publish(ctx context.Context, channel string, msg string) error {
	var err error
	log.Debug("Publish ...", msg)
	err = Red.Publish(ctx, channel, msg).Err()
	if err != nil {
		log.Error(err)
	}
	return err
}

//Subscribe 订阅Redis消息
func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := Red.Subscribe(ctx, channel)
	log.Debug("Subscribe ...", ctx)
	msg, err := sub.ReceiveMessage(ctx)
	if err != nil {
		log.Error(err)
		return "", err
	}
	log.Debug("Subscribe ...", msg.Payload)
	return msg.Payload, err
}
func GenMsgKey(userId,targetId int64)string{
	return fmt.Sprintf("msg_%d_%d",userId,targetId)
}


func MsgAddRedis(ctx context.Context,targetId int64,data any){
	key := fmt.Sprintf("%d",targetId)
	dataJSON, err := json.Marshal(data)
	if err != nil {
		log.Error("redis data marshal err",err)
	}
	dataStr := string(dataJSON)
	ress,err:=Red.SAdd(ctx,key,dataStr).Result()
	if err != nil {
		log.Error(err)
	}
	log.Debug(ress)
}