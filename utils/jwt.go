// Package utils
//
//	@Title			jwt.go
//	@Description	本文件为JWT令牌的管理
package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrEmptyToken   = errors.New("empty token")
	ErrInvalidToken = errors.New("invalid token")
)

// GenerateJWTToken 生成JWT令牌
func GenerateJWTToken(method jwt.SigningMethod, claims jwt.Claims, secret any) (string, error) {
	token := jwt.NewWithClaims(method, claims)
	return token.SignedString(secret)
}

// ParseJWTToken 解析JWT令牌
func ParseJWTToken(tokenSignedString string, claims jwt.Claims, secret any) (*jwt.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenSignedString, claims, func(token *jwt.Token) (any, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if token == nil {
		return nil, ErrEmptyToken
	}
	if !token.Valid {
		return nil, ErrInvalidToken
	}
	return &token.Claims, nil
}
