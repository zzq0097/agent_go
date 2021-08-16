package service

import (
	"agent/util"
	"os/exec"
	"strings"
)

type Status struct {
	Pid     string `json:"pid"`
	Running bool   `json:"running"`
	Exist   bool   `json:"exist"`
	Check   bool   `json:"check"`
	Version string `json:"version"`
}

func NgxStart(cmdFile string, confFile string) string {
	return execForStr(cmdFile + " -c " + confFile)
}

func NgxReload(cmdFile string, confFile string) string {
	return execForStr(cmdFile + " -s reload -c " + confFile)
}

func NgxStop(cmdFile string) string {
	return execForStr(cmdFile + " -s stop")
}

func NgxT(cmdFile string, confFile string) string {
	return execForStr(cmdFile + " -t -c " + confFile)
}

func NgxV(cmdFile string) string {
	return execForStr(cmdFile + " -V")
}

func NgxCheck(ngxT string) bool {
	return strings.HasSuffix(ngxT, "successful\n")
}

func NgxVersion(ngxV string) string {
	if strings.HasPrefix(ngxV, "nginx version: ") {
		return ngxV[0:strings.Index(ngxV, "\n")][15:]
	}
	return ""
}

func NgxExist(cmdFile string) bool {
	return util.FileExist(cmdFile)
}

func NgxStatus(cmdFile string, confFile string) interface{} {
	status := Status{}
	status.Pid = NgxPid()
	status.Running = NgxPid() != ""
	status.Exist = NgxExist(cmdFile)
	status.Version = NgxVersion(NgxV(cmdFile))
	status.Check = NgxCheck(NgxT(cmdFile, confFile))
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

func execForStr(command string) string {
	output, err := exec.Command("/bin/sh", "-c", command).CombinedOutput()
	if err != nil {
		return "error"
	}
	return string(output)
}
