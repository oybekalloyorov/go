package util

import "time"

// Backoff returns an exponential backoff duration for attempt (0-based).
// e.g. attempt 0 -> 500ms, attempt 1 -> 1s, attempt 2 -> 2s, attempt 3 -> 4s ...
func Backoff(attempt int) time.Duration {
	if attempt <= 0 {
		return 500 * time.Millisecond
	}
	return time.Duration(1<<uint(attempt)) * 500 * time.Millisecond
}
