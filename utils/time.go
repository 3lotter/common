// Package utils
//
//	@Title			time.go
//	@Description	本文件为time相关的方法
package utils

import "time"

const (
	DateFormat     string = "2006-01-02"
	DateTimeFormat string = "2006-01-02 15:04:05"
	UTCFormat      string = "2006-01-02T15:04:05Z"
)

// FormatTime 根据Time获取给定格式的本地格式化时间
func FormatTime(t time.Time, dateFormat string) string {
	return t.Format(dateFormat)
}

// FormatTimestamp 根据毫秒级时间戳获取按给定格式的本地格式化时间
func FormatTimestamp(timestamp int64, dateFormat string) string {
	t := time.UnixMilli(timestamp)
	return t.Format(dateFormat)
}

// ParseFormattedTime 按给定格式解析本地格式化时间
func ParseFormattedTime(datetime string, dateFormat string) (time.Time, error) {
	return time.Parse(dateFormat, datetime)
}
