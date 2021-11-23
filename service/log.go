package service

import (
	"errors"
	"os"
	"runtime"
)

func ReadLogs(path string, line int) ([]string, error) {
	var nl byte
	switch runtime.GOOS {
	case "linux":
		nl = '\n'
	case "darwin":
		nl = '\r'
	default:
		return nil, errors.New("only linux、darwin")
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	size := stat.Size()
	pos := size
	byt := make([]byte, 1)
	_, _ = file.ReadAt(byt, pos-1)
	if byt[0] == nl {
		pos--
	}
	l := 0
	lastLine := pos
	var arr []string
	for {
		if l >= line {
			break
		}
		if pos == 0 {
			byt = make([]byte, lastLine-1)
			_, err = file.ReadAt(byt, 0)
			arr = append(arr, string(byt))
			break
		}
		byt := make([]byte, 1)
		_, _ = file.ReadAt(byt, pos-1)
		if byt[0] == nl {
			byt = make([]byte, lastLine-pos)
			_, err = file.ReadAt(byt, pos)
			arr = append(arr, string(byt))
			lastLine = pos
			l++
		}
		pos--
	}
	return arr, nil
}
