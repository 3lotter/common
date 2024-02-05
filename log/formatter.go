// Package log
//	@Title			formatter.go
//	@Description	本文件定义了日志输出格式
package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

const (
	timestampFormat string = "2006-01-02 15:04:05" // yyyy-MM-dd hh-mm-ss
	dataLocation           = "location"            // 用于记录报错位置
)

// ServerFormatter 服务日志格式
type ServerFormatter struct{}

// Format 服务日志按自定义格式输出
// 格式为：|日志级别| 时间戳 (报错代码位置) 报错内容
func (serverFormatter *ServerFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format(timestampFormat)
	msg := fmt.Sprintf("%s |%s| %s\n", timestamp, strings.ToUpper(entry.Level.String()), entry.Message)

	if entry.Data != nil {
		if location, ok := entry.Data[dataLocation]; ok {
			msg = fmt.Sprintf("%s |%s| (%s) %s\n", timestamp, strings.ToUpper(entry.Level.String()), location.(string), entry.Message)
		}
	}
	return []byte(msg), nil
}
