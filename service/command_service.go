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
	ExecMsg string `json:"execMsg"`
}

func start(cmdFile string, confFile string) string {
	return execForStr(cmdFile + " -c " + confFile)
}

func reload(cmdFile string, confFile string) string {
	return execForStr(cmdFile + " -s reload -c " + confFile)
}

func stop(cmdFile string) string {
	return execForStr(cmdFile + " -s stop")
}

func check(cmdFile string, confFile string) string {
	return execForStr(cmdFile + " -t -c " + confFile)
}

func NgxStart(cmdFile string, confFile string) *Status {
	return NgxStatus(cmdFile, confFile, start(cmdFile, confFile))
}

func NgxReload(cmdFile string, confFile string) *Status {
	return NgxStatus(cmdFile, confFile, reload(cmdFile, confFile))
}

func NgxStop(cmdFile string, confFile string) *Status {
	return NgxStatus(cmdFile, confFile, stop(cmdFile))
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

// NgxStatus args 只会获取第一个参数
func NgxStatus(cmdFile string, confFile string, args ...string) *Status {
	pid := NgxPid()
	checkMsg := check(cmdFile, confFile)
	var execMsg string
	if len(args) == 1 {
		execMsg = args[0]
	} else {
		execMsg = checkMsg
	}
	return &Status{
		Pid:     pid,
		Running: pid != "",
		Exist:   NgxExist(cmdFile),
		Version: NgxVersion(NgxV(cmdFile)),
		Check:   NgxCheck(checkMsg),
		ExecMsg: execMsg,
	}
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
		return err.Error()
	}
	return string(output)
}
