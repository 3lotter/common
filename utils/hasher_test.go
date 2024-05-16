package utils

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"github.com/tjfoc/gmsm/sm3"
	"golang.org/x/crypto/ripemd160"
	"testing"
)

func TestSM3(t *testing.T) {
	input := []byte("hello world")
	expectedOutput := sm3.Sm3Sum(input)
	output, err := SM3(input)
	if err != nil {
		t.Errorf("SM3 returned an error: %v", err)
	}
	if !bytes.Equal(output, expectedOutput) {
		t.Errorf("SM3 output was incorrect, got: %x, want: %x", output, expectedOutput)
	}
}

func TestMD5(t *testing.T) {
	input := []byte("hello world")
	hasher := md5.New()
	hasher.Write(input)
	expectedOutput := hasher.Sum(nil)
	output, err := MD5(input)
	if err != nil {
		t.Errorf("MD5 returned an error: %v", err)
	}
	if !bytes.Equal(output, expectedOutput) {
		t.Errorf("MD5 output was incorrect, got: %x, want: %x", output, expectedOutput)
	}
}

func TestRIPEMD160(t *testing.T) {
	input := []byte("hello world")
	hasher := ripemd160.New()
	hasher.Write(input)
	expectedOutput := hasher.Sum(nil)
	output, err := RIPEMD160(input)
	if err != nil {
		t.Errorf("RIPEMD160 returned an error: %v", err)
	}
	if !bytes.Equal(output, expectedOutput) {
		t.Errorf("RIPEMD160 output was incorrect, got: %x, want: %x", output, expectedOutput)
	}
}

func TestSHA256(t *testing.T) {
	input := []byte("hello world")
	hasher := sha256.New()
	hasher.Write(input)
	expectedOutput := hasher.Sum(nil)
	output, err := SHA256(input)
	if err != nil {
		t.Errorf("SHA256 returned an error: %v", err)
	}
	if !bytes.Equal(output, expectedOutput) {
		t.Errorf("SHA256 output was incorrect, got: %x, want: %x", output, expectedOutput)
	}
}
