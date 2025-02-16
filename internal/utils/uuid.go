package utils

import (
	uuid "github.com/google/uuid"
)

// String returns string pointer
func String(data string) *string {
	return &data
}

// Int returns int pointer
func Int(data int) *int {
	return &data
}

// IntUnref returns zero if argument is nil, otherwise it returns int value.
func IntUnref(data *int) int {
	if data == nil {
		return 0
	}
	return *data
}

// StringUnref returns empty string if argument is nil, otherwise it returns string value.
func StringUnref(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// Bool creates reference from given data value.
func Bool(data bool) *bool {
	return &data
}

// BoolUnref returns false if argument is nil, otherwise it returns bool value.
func BoolUnref(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

// UUID returns UUID
func Uuid() string {
	id := uuid.New()
	return id.String()
}

// IsValidUUID validate
func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
