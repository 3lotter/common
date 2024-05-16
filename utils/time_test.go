package utils

import (
	"testing"
	"time"
)

func TestFormatTime(t *testing.T) {
	// Set up a specific time in UTC
	t1 := time.Date(2023, time.March, 25, 12, 34, 56, 0, time.UTC)
	// Convert UTC time to local time for the expected result
	localTime := t1.In(time.Local)
	expected := localTime.Format(DateTimeFormat) // Format the local time according to the DateTimeFormat

	// Call the function with the UTC time
	result := FormatTime(t1, DateTimeFormat)
	if result != expected {
		t.Errorf("FormatTime failed, expected %q, got %q", expected, result)
	}
}

func TestFormatTimestamp(t *testing.T) {
	// Positive test case
	timestamp := int64(1679760896000) // corresponds to 2023-03-25 12:34:56 UTC
	expected := time.UnixMilli(timestamp).In(time.Local).Format(DateTimeFormat)
	result := FormatTimestamp(timestamp, DateTimeFormat)
	if result != expected {
		t.Errorf("FormatTimestamp failed, expected %v, got %v", expected, result)
	}

	// Negative test case: using incorrect format
	result = FormatTimestamp(timestamp, "01-02-2006")
	if result == expected {
		t.Errorf("FormatTimestamp should fail when formatting with incorrect format")
	}
}

func TestParseFormattedTime(t *testing.T) {
	// Positive test case
	datetime := "2023-03-25 12:34:56"
	expected, _ := time.ParseInLocation(DateTimeFormat, datetime, time.Local)
	parsedTime, err := ParseFormattedTime(datetime, DateTimeFormat)
	if err != nil || !parsedTime.Equal(expected) {
		t.Errorf("ParseFormattedTime failed, expected %v, got %v", expected, parsedTime)
	}

	// Negative test case: incorrect format
	_, err = ParseFormattedTime(datetime, "01-02-2006")
	if err == nil {
		t.Errorf("ParseFormattedTime should fail when parsing with incorrect format")
	}
}
