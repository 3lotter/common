package utils

import (
	"testing"
)

func TestJsonCompare(t *testing.T) {
	// Define test cases
	tests := []struct {
		name     string
		jsonA    []byte
		jsonB    []byte
		expected bool
		wantErr  bool
	}{
		{
			name:     "both empty",
			jsonA:    []byte("{}"),
			jsonB:    []byte("{}"),
			expected: true,
			wantErr:  false,
		},
		{
			name:     "identical json objects",
			jsonA:    []byte(`{"name":"John", "age":30}`),
			jsonB:    []byte(`{"name":"John", "age":30}`),
			expected: true,
			wantErr:  false,
		},
		{
			name:     "different json objects",
			jsonA:    []byte(`{"name":"John", "age":30}`),
			jsonB:    []byte(`{"name":"Jane", "age":25}`),
			expected: false,
			wantErr:  false,
		},
		{
			name:     "invalid json in jsonA",
			jsonA:    []byte(`{"name": "John", "age":30`),
			jsonB:    []byte(`{"name":"John", "age":30}`),
			expected: false,
			wantErr:  true,
		},
		{
			name:     "invalid json in jsonB",
			jsonA:    []byte(`{"name":"John", "age":30}`),
			jsonB:    []byte(`{"age":30`),
			expected: false,
			wantErr:  true,
		},
		{
			name:     "nested json objects",
			jsonA:    []byte(`{"person": {"name":"John", "age":30}}`),
			jsonB:    []byte(`{"person": {"name":"John", "age":30}}`),
			expected: true,
			wantErr:  false,
		},
		{
			name:     "nested different json objects",
			jsonA:    []byte(`{"person": {"name":"John", "age":30}}`),
			jsonB:    []byte(`{"person": {"name":"Jane", "age":25}}`),
			expected: false,
			wantErr:  false,
		},
	}

	// Execute each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JsonCompare(tt.jsonA, tt.jsonB)
			if (err != nil) != tt.wantErr {
				t.Errorf("JsonCompare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("JsonCompare() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
