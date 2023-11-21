package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//配置文件使用Viper，初始化后，信息反序列化存到全局变量里

// 定义全局变量 服务配置
var Conf = new(SrvConfig)

// 或viper.Getxxx()读取

//Viper 使用的是`mapstructure`

type SrvConfig struct {
	Name    string `mapstructure:"name"`
	Mode    string `mapstructure:"mode"`
	Version string `mapstructure:"version"`
	// snowflake用
	StartTime string `mapstructure:"start_time"`
	MachineId int64  `mapstructure:"machine_id"`
	Udp       int64  `mapstructure:"udp"`

	IP       string `mapstructure:"ip"`
	HttpPort int    `mapstructure:"httpPort"`

	Services      map[string]*Service `mapstructure:"services"`
	*Etcd         `mapstructure:"etcd"`
	*LogConfig    `mapstructure:"log"`
	*MySQLConfig  `mapstructure:"mysql"`
	*RedisConfig  `mapstructure:"redis"`
	*ConsulConfig `mapstructure:"consul"`
}

type Service struct {
	Name        string `mapstructure:"name"`
	RpcPort     int    `mapstructure:"rpcPort"`
	LoadBalance bool   `mapstructure:"loadBalance"`
}

type Etcd struct {
	Endpoints   []string `mapstructure:"endpoints"`
	DialTimeout int      `mapstructure:"dialTimeout"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"dbname"`
	Auto         bool   `mapstructure:"auto"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Password     string `mapstructure:"password"`
	Port         int    `mapstructure:"port"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns"`
}

type ConsulConfig struct {
	Addr string `mapstructure:"addr"`
}

// 所有配置文件初始化
func Init(filePath string) (err error) {
	//方式1：直接指定配置文件路径（相对路径或者绝对路径）
	//相对路径：相对执行的可执行文件的相对路径
	// viper.SetConfigFile("./conf/config.yaml")
	if filePath == "" {
		filePath = "config/config.yaml"
	}
	viper.SetConfigFile(filePath)

	err = viper.ReadInConfig() //读取配置信息
	if err != nil {
		//读取配置信息失败
		fmt.Printf("viper.ReadInConfig failed,err:%v\n", err)
	}
	//如果使用viper.get,就不用下一步

	//把读取到的配置信息反序列化到Conf变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed,err:%v\n", err)
	}

	viper.WatchConfig() //配置文件监听
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了。。")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed,err: %v\n", err)
		}
	})
	return
}
