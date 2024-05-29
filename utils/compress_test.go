package utils

import (
	"testing"
)

func TestCompressString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"EmptyString", "", false},
		{"SimpleString", "Hello, World!", false},
		{"LongString", "This is a longer string that we want to compress and see if it works correctly.", false},
		{"SpecialCharacters", "!@#$%^&*()_+=-[]{}|;:,.<>?/~`", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			compressed, err := CompressString(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("CompressString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				// Decompress to check if the original string is obtained
				decompressed, err := DeCompressString(compressed)
				if err != nil {
					t.Errorf("DeCompressString() error = %v", err)
					return
				}
				if decompressed != tt.input {
					t.Errorf("DeCompressString() = %v, want %v", decompressed, tt.input)
				}
			}
		})
	}
}

func TestDeCompressString(t *testing.T) {
	// 先生成一个有效的压缩字符串
	compressed, err := CompressString("Hello, World!")
	if err != nil {
		t.Fatalf("Failed to compress string: %v", err)
	}

	tests := []struct {
		name        string
		input       string
		wantErr     bool
		expectedOut string
	}{
		{"EmptyString", "", true, ""},
		{"InvalidBase64", "InvalidBase64String", true, ""},
		{"ValidCompressedString", compressed, false, "Hello, World!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decompressed, err := DeCompressString(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeCompressString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && decompressed != tt.expectedOut {
				t.Errorf("DeCompressString() = %v, want %v", decompressed, tt.expectedOut)
			}
		})
	}
}
