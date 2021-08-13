package service

import "os/exec"

func NgxStart(cmdFile string, confFile string) string {
	exec.Command(cmdFile, "-c", confFile)
	return ""
}

func NgxReload(cmdFile string, confFile string) string {
	exec.Command(cmdFile, "-s", "reload", "-c", confFile)
	return ""
}

func NgxStop(cmdFile string) string {
	exec.Command(cmdFile, "-s", "stop")
	return ""
}

func NgxCheck(cmdFile string, confFile string) string {
	exec.Command(cmdFile, "-t", "-c", confFile)
	return ""
}

func NgxPid() {
	exec.Command("/bin/sh", "-c", "ps -ef|grep nginx")
}

func NgxStatus() {

}

func NgxVersion(cmdFile string) string {
	exec.Command(cmdFile, "-V")
	return ""
}

func execForStr() string {
	exec.Command("", "", "")
	return ""
}
