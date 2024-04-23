// Package utils
//
//	@Title			hasher.go
//	@Description	本文件为密码学相关的方法
package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"github.com/tjfoc/gmsm/sm3"
	"golang.org/x/crypto/ripemd160"
)

func SM3(input []byte) ([]byte, error) {
	hasher := sm3.New()
	if _, err := hasher.Write(input); err != nil {
		return nil, err
	}
	return hasher.Sum(nil), nil
}

func MD5(input []byte) ([]byte, error) {
	hasher := md5.New()
	if _, err := hasher.Write(input); err != nil {
		return nil, err
	}
	return hasher.Sum(nil), nil
}

func RIPEMD160(input []byte) ([]byte, error) {
	hasher := ripemd160.New()
	if _, err := hasher.Write(input); err != nil {

		return nil, err
	}
	return hasher.Sum(nil), nil
}

func SHA256(input []byte) ([]byte, error) {
	hasher := sha256.New()
	if _, err := hasher.Write(input); err != nil {
		return nil, err
	}
	return hasher.Sum(nil), nil
}
