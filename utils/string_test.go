package utils

import (
	"testing"
)

func TestConcat(t *testing.T) {
	// Positive test cases
	t.Run("Multiple strings", func(t *testing.T) {
		result := Concat("Hello", " ", "world", "!")
		expected := "Hello world!"
		if result != expected {
			t.Errorf("Concat failed: Expected %q, got %q", expected, result)
		}
	})

	t.Run("Two strings", func(t *testing.T) {
		result := Concat("Go", "Lang")
		expected := "GoLang"
		if result != expected {
			t.Errorf("Concat failed: Expected %q, got %q", expected, result)
		}
	})

	t.Run("With empty string", func(t *testing.T) {
		result := Concat("Hello", "", "World")
		expected := "HelloWorld"
		if result != expected {
			t.Errorf("Concat failed: Expected %q, got %q", expected, result)
		}
	})

	t.Run("All empty strings", func(t *testing.T) {
		result := Concat("", "", "")
		expected := ""
		if result != expected {
			t.Errorf("Concat failed: Expected %q, got %q", expected, result)
		}
	})
}
