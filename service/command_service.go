package service

import (
	"os"
	"os/exec"
	"strings"
)

func NgxStart(cmdFile string, confFile string) string {
	return execForStr(cmdFile + " -c " + confFile)
}

func NgxReload(cmdFile string, confFile string) string {
	return execForStr(cmdFile + " -s reload -c " + confFile)
}

func NgxStop(cmdFile string) string {
	return execForStr(cmdFile + " -s stop")
}

func NgxCheck(cmdFile string, confFile string) string {
	return execForStr(cmdFile + " -t -c " + confFile)
}

func NgxV(cmdFile string) string {
	return execForStr(cmdFile + " -V")
}

func NgxStatus(cmdFile string, confFile string) interface{} {
	type Status struct {
		Pid     string `json:"pid"`
		Running bool   `json:"running"`
		Exist   bool   `json:"exist"`
		Check   bool   `json:"check"`
		Version string `json:"version"`
	}
	status := Status{}
	status.Pid = NgxPid()
	status.Running = NgxPid() != ""
	status.Exist = FileExist(cmdFile)
	status.Version = NgxVersion(cmdFile)
	status.Check = strings.HasSuffix(NgxCheck(cmdFile, confFile), "successful\n")
	return status
}

func NgxPid() string {
	pid := ""
	pidLines := execForStr("ps -ef|grep nginx")
	for _, pidLine := range strings.Split(pidLines, "\n") {
		if strings.Contains(pidLine, "nginx: master process") {
			first := true
			for _, s := range strings.Split(pidLine, " ") {
				if s == "" {
					continue
				}
				if first {
					first = false
					continue
				}
				pid = s
				break
			}
			break
		}
	}
	return pid
}

func NgxVersion(cmdFile string) string {
	lines := execForStr(cmdFile + " -V")
	if strings.HasPrefix(lines, "nginx version: ") {
		line := lines[0:strings.Index(lines, "\n")]
		return line[15:]
	}
	return ""
}

func execForStr(command string) string {
	output, err := exec.Command("/bin/sh", "-c", command).CombinedOutput()
	if err != nil {
		return "error"
	}
	return string(output)
}

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
