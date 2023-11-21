package mysql

import (
	"fmt"
	"time"

	"github.com/disgoorg/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"one_ser.com/app/chat_service/sys_init/model"
	"one_ser.com/config"
)

var DB *gorm.DB
type gormWriter	struct{}
func (w gormWriter) Printf(format string, args ...interface{}) {
	log.Infof("format",args...)
}

func Init(cfg *config.MySQLConfig) (err error) {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	// host:prot@tcp(username:password)/dbname?编码&time类型转换&所在地
	newLogger := logger.New(
		gormWriter{},
		logger.Config{
			SlowThreshold: 500 * time.Millisecond,
			LogLevel: logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful: false,
		},

	)


	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return err
	}
	if cfg.Auto {
		DB.AutoMigrate(&model.Message{})
	}

	// 额外的连接配置，拿取原生db
	sqlDB, err := DB.DB() // database/sql.DB
	if err != nil {
		return
	}

	// 以下配置要配合 mysql.conf 进行配置
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	return
}
