package crawler

import "time"

// getSleepDuration returns the sleep duration based on the first element of sleepMs.
// If sleepMs is empty, it returns the default duration of 1000 milliseconds.
func getSleepDuration(sleepMs ...int) time.Duration {
	// Use default sleep duration of 1000 milliseconds
	duration := 1000 * time.Millisecond

	// If a custom sleep duration is provided, use the first value
	if len(sleepMs) > 0 {
		duration = time.Duration(sleepMs[0]) * time.Millisecond
	}

	return duration
}
