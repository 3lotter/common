// Package databases
//
//	@Title			databases.go
//	@Description	本文件用于配置数据库连接，支持mysql、pgsql和达梦数据库
package databases

import (
	"database/sql"
	dm "github.com/3lotter/dm-gorm2-dialect"
	xugu "github.com/3lotter/xugu-gorm2-dialect"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

const (
	mysqlDBMS       string = "mysql"
	postgresqlDBMS  string = "pgsql"
	dmsqlDBMS       string = "dmsql"
	xugusqlDBMS     string = "xugusql"
	xuguclusterDBMS string = "xugucluster"
)

// 各数据库连接语句默认模板
const (
	mysqlDsnFormat       string = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local%s"
	postgresqlDsnFormat  string = "postgres://%s:%s@%s:%d/%s?sslmode=disable%s"
	dmsqlDsnFormat       string = "dm://%s:%s@%s:%d?schema=%s%s"
	xugusqlDsnFormat     string = "User=%s;PWD=%s;IP=%s;Port=%d;DB=%s;CURRENT_SCHEMA=%s%s"
	xuguclusterDsnFormat string = "User=%s;PWD=%s;IPS=%s;Port=%d;DB=%s;CURRENT_SCHEMA=%s%s"
)

var DsnMap = map[string]string{
	mysqlDBMS:       mysqlDsnFormat,
	postgresqlDBMS:  postgresqlDsnFormat,
	dmsqlDBMS:       dmsqlDsnFormat,
	xugusqlDBMS:     xugusqlDsnFormat,
	xuguclusterDBMS: xuguclusterDsnFormat,
}

type Conf struct {
	DBMS            string          // 数据库管理系统名称，支持 [mysql|postgresql|dmsql|xugusql]
	Username        string          // 登陆用户名
	Password        string          // 登陆密码
	IP              string          // 数据库IP
	Port            int             // 数据库端口
	DBName          string          // 数据库名称
	Schema          string          // 模式名称
	External        string          // 扩展内容
	DSN             string          // 数据库链接字符串
	MaxIdleConn     int             // 最大空闲连接数
	MaxOpenConn     int             // 最大连接数
	ConnMaxLifeTime int             // 连接最大生命时间
	LogMode         logger.LogLevel // 日志级别 [1:silent|2:error|3:warn|4:info]
	IsSingularTable bool            // 表名是否要求单数
}

func (conf Conf) GenerateDSN() (string, error) {
	dsnFormat := DsnMap[conf.DBMS]

	switch conf.DBMS {
	case mysqlDBMS:
		return fmt.Sprintf(dsnFormat, conf.Username, conf.Password, conf.IP, conf.Port, conf.DBName, conf.External), nil
	case postgresqlDBMS, dmsqlDBMS:
		return fmt.Sprintf(dsnFormat, conf.Username, conf.Password, conf.IP, conf.Port, conf.Schema, conf.External), nil
	case xugusqlDBMS, xuguclusterDBMS:
		return fmt.Sprintf(dsnFormat, conf.Username, conf.Password, conf.IP, conf.Port, conf.DBName, conf.Schema, conf.External), nil
	}
	return "", fmt.Errorf("DBMS [%s] not supported", conf.DBMS)
}

type MyGormDB struct {
	*gorm.DB
}

// SetupDB 初始化DB连接
func SetupDB(conf Conf) (*MyGormDB, error) {
	var err error

	// 如果dsn为空，则用通用模板生成dsn
	dbms := conf.DBMS
	if len(conf.DSN) == 0 {
		conf.DSN, err = conf.GenerateDSN()
		if err != nil {
			return nil, errors.WithMessage(err, "conf.GenerateDSN fail")
		}
	}

	// 生成gorm配置
	gormConf := &gorm.Config{
		Logger: logger.Default.LogMode(conf.LogMode),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: conf.IsSingularTable,
		},
	}

	// 建立连接
	var gormDB *gorm.DB
	switch dbms {
	case mysqlDBMS:
		gormDB, err = gorm.Open(mysql.Open(conf.DSN), gormConf)
	case postgresqlDBMS:
		gormDB, err = gorm.Open(postgres.Open(conf.DSN), gormConf)
	case dmsqlDBMS:
		gormDB, err = gorm.Open(dm.Open(conf.DSN), gormConf)
	case xugusqlDBMS, xuguclusterDBMS:
		gormDB, err = gorm.Open(xugu.Open(conf.DSN), gormConf)
	default:
		return nil, errors.New(fmt.Sprintf("DBMS [%s] not supported", dbms))
	}
	if err != nil {
		return nil, errors.Wrapf(err, "gorm.Open fail, dsn: %s", conf.DSN)
	}

	// 连通性测试
	var sqlDb *sql.DB
	sqlDb, err = gormDB.DB()
	if err != nil {
		return nil, errors.Wrap(err, "gormDB.DB fail")
	}
	err = sqlDb.Ping()
	if err != nil {
		sqlDb.Close()
		return nil, errors.Wrap(err, "sqlDb.Ping fail")
	}

	// 设置连接数
	sqlDb.SetMaxIdleConns(conf.MaxIdleConn)
	sqlDb.SetMaxOpenConns(conf.MaxOpenConn)
	sqlDb.SetConnMaxLifetime(time.Duration(conf.ConnMaxLifeTime) * time.Second)

	return &MyGormDB{DB: gormDB}, nil
}

// CloseDB 关闭数据库连接
func (myGormDB *MyGormDB) CloseDB() error {
	db, err := myGormDB.DB.DB()
	if err != nil {
		return errors.Wrap(err, "gormDB.DB fail")
	}
	return db.Close()
}
