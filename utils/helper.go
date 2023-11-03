package utils

import "os"

func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func FormatArrayStartNEndIdx(start, stop, n int) (int, int) {
	if start < 0 {
		start = 0
	}
	if stop < 0 {
		stop = n + stop
	}
	if stop >= n {
		stop = n - 1
	}
	return start, stop
}
