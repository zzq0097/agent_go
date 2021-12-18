package service

import (
	"runtime"
	"strconv"
	"strings"
)

type PortUsed struct {
	NetType string `json:"netType"`
	State   string `json:"state"`
	Port    int    `json:"port"`
	Name    string `json:"name"`
}

// CheckPortUsed
// example：tcp LISTEN 0 128 0.0.0.0:22 0.0.0.0:* users:(("sshd",pid=986,fd=5))
func CheckPortUsed() []PortUsed {
	str := ExecForStr("ss -tunlp")
	lines := strings.Split(str, "\n")
	var arr []PortUsed
	for i := 1; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		col := strings.Split(lines[i], " ")
		idx := 0
		var portUsed PortUsed
		for _, c := range col {
			if c == "" {
				continue
			} else {
				switch idx {
				case 0:
					portUsed.NetType = c
				case 1:
					portUsed.State = c
				case 4:
					port, _ := strconv.Atoi(c[strings.LastIndex(c, ":")+1:])
					portUsed.Port = port
				case 6:
					trimLeft := strings.Replace(c, "users:((\"", "", 1)
					portUsed.Name = trimLeft[:strings.Index(trimLeft, "\"")]
				}
				idx++
			}
		}
		arr = append(arr, portUsed)
	}
	return arr
}

func Os() string {
	os := runtime.GOOS
	if os == "linux" {

	}
	return os
}
