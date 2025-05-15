package crawler

import "time"

// getSleepDuration returns the sleep duration based on the first element of sleepMs.
// If sleepMs is empty, it returns the default duration of 1000 milliseconds.
//
// Parameters:
//   - sleepMs: Optional variadic integer slice representing the sleep duration in milliseconds.
//
// Returns:
//   - time.Duration: The calculated sleep duration.
//
// Example:
//
//	getSleepDuration()        // returns 1000ms
//	getSleepDuration(500)     // returns 500ms
func getSleepDuration(sleepMs ...int) time.Duration {
	// Use default sleep duration of 1000 milliseconds
	duration := 1000 * time.Millisecond

	// If a custom sleep duration is provided, use the first value
	if len(sleepMs) > 0 {
		duration = time.Duration(sleepMs[0]) * time.Millisecond
	}

	return duration
}
