// Package databases
//
//	@Title			database.go
//	@Description	本文件用于配置数据库连接，支持mysql、pgsql和达梦数据库
package databases

import (
	dm "codeup.aliyun.com/6308f33e9011ed4f984a7e9d/dm-gorm2-dialect"
	"database/sql"
	"gorm.io/gorm/logger"

	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

// Conf DB连接配置
type Conf struct {
	DBMS            string
	DSN             string
	MaxIdleConn     int
	MaxOpenConn     int
	ConnMaxLifeTime int
	LogMode         logger.LogLevel
}

type MyGormDB struct {
	*gorm.DB
}

// SetupDB 初始化DB连接
func SetupDB(conf Conf) (*MyGormDB, error) {
	var err error

	// 建立连接
	var gormDb *gorm.DB
	dbms := conf.DBMS

	gormConf := &gorm.Config{
		Logger: logger.Default.LogMode(conf.LogMode),
	}

	switch dbms {
	case "mysql":
		gormDb, err = gorm.Open(mysql.Open(conf.DSN), gormConf)
	case "pgsql":
		gormDb, err = gorm.Open(postgres.Open(conf.DSN), gormConf)
	case "dmsql":
		gormDb, err = gorm.Open(dm.Open(conf.DSN), gormConf)
	default:
		return nil, errors.New(fmt.Sprintf("DBMS [%s] not supported", dbms))
	}
	if err != nil {
		return nil, err
	}

	// 连通性测试
	var sqlDb *sql.DB
	sqlDb, err = gormDb.DB()
	if err != nil {
		return nil, err
	}
	err = sqlDb.Ping()
	if err != nil {
		return nil, err
	}

	// 设置连接数
	sqlDb.SetMaxIdleConns(conf.MaxIdleConn)
	sqlDb.SetMaxOpenConns(conf.MaxOpenConn)
	sqlDb.SetConnMaxLifetime(time.Duration(conf.ConnMaxLifeTime) * time.Second)

	return &MyGormDB{DB: gormDb}, nil
}

// CloseDB 关闭数据库连接
func (myGormDB *MyGormDB) CloseDB() error {
	db, err := myGormDB.DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
