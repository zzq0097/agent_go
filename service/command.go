package service

import (
	"agent/util"
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

type Status struct {
	Pid     string `json:"pid"`
	Running bool   `json:"running"`
	Exist   bool   `json:"exist"`
	Check   bool   `json:"check"`
	Version string `json:"version"`
	CfgArgs string `json:"cfgArgs"`
	ExecMsg string `json:"execMsg"`
}

type ConfFile struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func start(cmdFile string, confFile string) string {
	return ExecForStr(cmdFile + " -c " + confFile)
}

func reload(cmdFile string, confFile string) string {
	return ExecForStr(cmdFile + " -s reload -c " + confFile)
}

func stop(cmdFile string) string {
	return ExecForStr(cmdFile + " -s stop")
}

func check(cmdFile string, confFile string) string {
	return ExecForStr(cmdFile + " -t -c " + confFile)
}

func ngxT(cmdFile string, confFile string) string {
	return ExecForStr(cmdFile + " -T -c " + confFile)
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
	return ExecForStr(cmdFile + " -V")
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

func NgxCfgArgs(ngxV string) string {
	if strings.Contains(ngxV, "configure arguments: ") {
		return strings.TrimRight(ngxV[strings.Index(ngxV, "configure arguments: ")+21:], "\n")
	}
	return ""
}

func NgxTConf(cmdFile string, confFile string) []ConfFile {
	var confFiles []ConfFile
	var content string
	var name string
	lines := strings.Split(ngxT(cmdFile, confFile), "\n")
	for i := range lines {
		if strings.HasPrefix(lines[i], "#") {
			if name != "" {
				confFiles = append(confFiles, ConfFile{name, content})
			}
			name = strings.TrimRight(lines[i][21:], ":")
			content = ""
		} else {
			content += lines[i] + "\n"
		}
	}
	return confFiles
}

func NgxExist(cmdFile string) bool {
	return util.FileExist(cmdFile)
}

// NgxStatus args 只会获取第一个参数
func NgxStatus(cmdFile string, confFile string, args ...string) *Status {
	pid := NgxPid()
	checkMsg := check(cmdFile, confFile)
	ngxV := NgxV(cmdFile)
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
		Version: NgxVersion(ngxV),
		CfgArgs: NgxCfgArgs(ngxV),
		Check:   NgxCheck(checkMsg),
		ExecMsg: execMsg,
	}
}

func NgxPid() string {
	pid := ""
	pidLines := ExecForStr("ps -ef|grep nginx")
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

func ExecForStr(command string) string {
	fmt.Println("[" + command + "]\n")
	var msg string
	cmd := exec.Command("/bin/sh", "-c", command)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err.Error()
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err.Error()
	}

	err = cmd.Start()
	if err != nil {
		return err.Error()
	}

	buf := 128
	reader := bufio.NewReader(stdout)
	for {
		byt := make([]byte, buf)
		i, err := reader.Read(byt)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err.Error()
		}
		msg += string(byt[0:i])
	}

	reader = bufio.NewReader(stderr)
	for {
		byt := make([]byte, buf)
		i, err := reader.Read(byt)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err.Error()
		}
		msg += string(byt[0:i])
	}

	_ = cmd.Wait()
	fmt.Println(msg)
	return msg
}
