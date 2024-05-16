package utils

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestGenerateJWTToken(t *testing.T) {
	// Define test claims
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		Issuer:    "test",
	}

	// Test positive case
	tokenString, err := GenerateJWTToken(jwt.SigningMethodHS256, claims, []byte("secret"))
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	// Test negative case with invalid secret
	_, err = GenerateJWTToken(jwt.SigningMethodHS256, claims, nil)
	assert.Error(t, err)
}

func TestParseJWTToken(t *testing.T) {
	// Define test claims and token
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		Issuer:    "test",
	}
	tokenString, _ := GenerateJWTToken(jwt.SigningMethodHS256, claims, []byte("secret"))

	// Test positive case
	parsedClaims := &jwt.RegisteredClaims{}
	_, err := ParseJWTToken(tokenString, parsedClaims, []byte("secret"))
	assert.NoError(t, err)
	assert.Equal(t, "test", parsedClaims.Issuer)

	// Test negative case with invalid token
	_, err = ParseJWTToken("invalidtoken", parsedClaims, []byte("secret"))
	assert.Error(t, err)

	// Test negative case with wrong secret
	_, err = ParseJWTToken(tokenString, parsedClaims, []byte("wrongsecret"))
	assert.Error(t, err)

	// Test token expiration
	expiredClaims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(-time.Hour)), // token expired 1 hour ago
		Issuer:    "test",
	}
	expiredTokenString, _ := GenerateJWTToken(jwt.SigningMethodHS256, expiredClaims, []byte("secret"))
	_, err = ParseJWTToken(expiredTokenString, expiredClaims, []byte("secret"))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "token is expired") // Check if the error message contains "token is expired"
}
