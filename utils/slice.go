// Package utils
//
//	@Title			slice.go
//	@Description	本文件为slice相关的方法
package utils

// Reverse 翻转切片
func Reverse[T any](input []T) {
	length := len(input)
	for i := 0; i < length/2; i++ {
		input[i], input[length-1-i] = input[length-1-i], input[i]
	}
}
