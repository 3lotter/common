package databases

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestGenerateDSN(t *testing.T) {
	conf := Conf{
		DBMS:     "mysql",
		Username: "root",
		Password: "password",
		IP:       "localhost",
		Port:     3306,
		DBName:   "test_db",
		Schema:   "",
		External: "",
		DSN:      "",
	}
	dsn, err := conf.GenerateDSN()
	assert.NoError(t, err)
	assert.NotEmpty(t, dsn)

	conf.DBMS = "unknown"
	_, err = conf.GenerateDSN()
	assert.Error(t, err)
}

func TestCloseDB(t *testing.T) {
	// 创建sqlmock对象
	sqlDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer sqlDB.Close()

	// 配置预期行为：期望数据库的Close方法被调用一次
	mock.ExpectClose()

	// 使用gorm和sqlmock创建模拟的gorm.DB对象
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	assert.NoError(t, err)

	// 创建MyGormDB实例
	myGormDB := &MyGormDB{DB: gormDB}

	// 调用CloseDB并断言没有错误返回
	err = myGormDB.CloseDB()
	assert.NoError(t, err)

	// 检查是否所有预期的数据库操作都已满足
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
