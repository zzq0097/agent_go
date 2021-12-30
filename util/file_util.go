package util

import "os"

func FileExist(fullPath string) bool {
	_, err := os.Stat(fullPath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func FileSize(fullPath string) int64 {
	stat, err := os.Stat(fullPath)
	if err != nil {
		return 0
	}
	return stat.Size()
}
