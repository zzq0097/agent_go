package service

import (
	"io/ioutil"
)

type ConfDFile struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type NgxConf struct {
	NginxConf string      `json:"nginx.conf"`
	ConfD     []ConfDFile `json:"conf.d"`
}

func AllConf(confDir string) NgxConf {
	mainFile := confDir + "/nginx.conf"
	confD := confDir + "/conf.d"
	file, err := ioutil.ReadFile(mainFile)
	if err != nil {
		return NgxConf{}
	}

	confFiles := make([]ConfDFile, 0)
	if dir, err := ioutil.ReadDir(confD); err == nil {
		for _, fs := range dir {
			filePath := confD + "/" + fs.Name()
			readFile, err := ioutil.ReadFile(filePath)
			if err != nil {
				continue
			}
			confFiles = append(confFiles, ConfDFile{Name: fs.Name(), Content: string(readFile)})
		}
	}
	return NgxConf{NginxConf: string(file), ConfD: confFiles}
}
