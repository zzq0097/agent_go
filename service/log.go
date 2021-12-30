package service

import (
	"agent/model"
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strconv"
)

func ReadReverse(path string, line int) ([]string, error) {
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
	defer file.Close()
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

func ReadOffset(path string, offset int64, sectionLog *model.SectionLog) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return
	}
	stat, err := file.Stat()
	if err != nil {
		return
	}

	size := stat.Size()
	pos := offset
	if size != 0 && pos < size {
		_, err = file.Seek(pos, 0)
		if err != nil {
			fmt.Println(err)
			return
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			LogLine(line, sectionLog)
		}
	}
}

func LogLine(line string, sectionLog *model.SectionLog) {
	log := model.AccessLog{}
	if err := json.Unmarshal([]byte(line), &log); err != nil {
		return
	}

	if reqLen, err := strconv.Atoi(log.RequestLength); err == nil {
		sectionLog.ReqLen += reqLen
	}

	if reqTime, err := strconv.ParseFloat(log.RequestTime, 32); err == nil {
		sectionLog.ReqTime += reqTime
	}
	if upsRespTime, err := strconv.ParseFloat(log.UpstreamResponseTime, 32); err == nil {
		sectionLog.UpsRespTime += upsRespTime
	}

	code := log.Status
	if _, ok := sectionLog.Code[code]; ok {
		sectionLog.Code[code]++
	} else {
		sectionLog.Code = make(map[string]int)
		sectionLog.Code[code] = 1
	}
	upsCode := log.UpstreamStatus
	if _, ok := sectionLog.UpsCode[upsCode]; ok {
		sectionLog.UpsCode[upsCode]++
	} else {
		sectionLog.UpsCode = make(map[string]int)
		sectionLog.UpsCode[upsCode] = 1
	}
}
