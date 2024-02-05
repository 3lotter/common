// Package utils
//
//	@Title			crypto.go
//	@Description	本文件为密码学相关的方法
package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
	"math/big"
)

func MD5(input []byte) ([]byte, error) {
	hasher := md5.New()
	_, err := hasher.Write(input)
	if err != nil {
		return nil, err
	}
	return hasher.Sum(nil), nil
}

func RIPEMD160(input []byte) ([]byte, error) {
	hasher := ripemd160.New()
	_, err := hasher.Write(input)
	if err != nil {
		return nil, err
	}
	return hasher.Sum(nil), nil
}

func SHA256(input []byte) ([]byte, error) {
	hasher := sha256.New()
	_, err := hasher.Write(input)
	if err != nil {
		return nil, err
	}
	return hasher.Sum(nil), nil
}

func Base58Encode(input []byte) []byte {
	var base58Alphabets = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
	x := big.NewInt(0).SetBytes(input)
	zero := big.NewInt(0)
	base, mod := big.NewInt(58), big.NewInt(0)

	var res []byte
	for x.Cmp(zero) > 0 {
		x.DivMod(x, base, mod)
		res = append(res, base58Alphabets[mod.Int64()])
	}

	Reverse[byte](res)
	return res
}
