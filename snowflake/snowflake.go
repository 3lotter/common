// Package snowflake
//
//	@Title			snowflake.go
//	@Description	本文件用于配置雪花编号生成器
package snowflake

import (
	"github.com/sony/sonyflake"
	"time"
)

type MySnowflakeGenerator struct {
	*sonyflake.Sonyflake
}

// SetupSnowflake 初始化snowflake生成器
func SetupSnowflake() (*MySnowflakeGenerator, error) {
	snowflakeGenerator, err := sonyflake.New(sonyflake.Settings{
		StartTime:      time.Time{},
		MachineID:      nil,
		CheckMachineID: nil,
	})
	if err != nil {
		return nil, err
	}
	return &MySnowflakeGenerator{snowflakeGenerator}, nil
}
