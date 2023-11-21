package main

import (
	"fmt"

	routes "one_ser.com/app/gateway/routers"
	"one_ser.com/app/gateway/rpc"
	"one_ser.com/utils"

	// "one_ser.com/app/gateway/sys_init/mysql"
	// "one_ser.com/app/gateway/sys_init/redis"
	"one_ser.com/config"
	"one_ser.com/logger"
)

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
	//3.初始化snowflake算法
	err = utils.SFInit(config.Conf.StartTime, config.Conf.MachineId)
	if err != nil {
		panic(err)
	}
	// 4. 初始rpc
	rpc.Init()

	//5. router
	r := routes.InitRouter()
	r.Run(fmt.Sprintf("%s:%d", config.Conf.IP, config.Conf.HttpPort))

}
