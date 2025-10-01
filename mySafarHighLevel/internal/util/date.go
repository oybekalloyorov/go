package util

import (
	"fmt"
	"time"
)

// ValidateDate parses dd.MM.yyyy and returns time.Time or error.
func ValidateDate(d string) (time.Time, error) {
	if d == "" {
		return time.Time{}, fmt.Errorf("empty date")
	}
	t, err := time.Parse("02.01.2006", d)
	if err != nil {
		return time.Time{}, fmt.Errorf("wrong format, want dd.MM.yyyy: %w", err)
	}
	return t, nil
}
