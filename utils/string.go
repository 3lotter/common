// Package utils
//
//	@Title			string.go
//	@Description	本文件为string相关的方法
package utils

import "strings"

// Concat 拼接字符串
func Concat(strs ...string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
	}
	return sb.String()
}
