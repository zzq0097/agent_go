package util

import "time"

const (
	Layout = "2006-01-02 15:04:05"
)

func Now() string {
	return time.Now().Format(Layout)
}
