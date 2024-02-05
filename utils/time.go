// Package utils
//
//	@Title			time.go
//	@Description	本文件为time相关的方法
package utils

import "time"

const (
	UTCFormat string = "2006-01-02T15:04:05Z"
)

// LocalUTC 获取UTC格式的本地时间
func LocalUTC(t time.Time) (string, error) {
	local, err := time.LoadLocation("Local")
	if err != nil {
		return "", err
	}
	return t.In(local).Format(UTCFormat), nil
}

// TimestampToUTC 解析时间戳，获取UTC格式的本地时间
func TimestampToUTC(timestamp int64) (string, error) {
	t := time.Unix(timestamp, 0)
	local, err := time.LoadLocation("Local")
	if err != nil {
		return "", err
	}
	return t.In(local).Format(UTCFormat), nil
}

// ParseUTC 解析UTC日期
func ParseUTC(datetime string) (time.Time, error) {
	return time.ParseInLocation(UTCFormat, datetime, time.Local)
}
