package service

import (
	"io/ioutil"
)

type ConfDFile struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type NgxConfD struct {
	NginxConf string      `json:"nginx.conf"`
	ConfD     []ConfDFile `json:"conf.d"`
}

func AllConfD(confDir string) NgxConfD {
	mainFile := confDir + "/nginx.conf"
	confD := confDir + "/conf.d"
	file, err := ioutil.ReadFile(mainFile)
	if err != nil {
		return NgxConfD{}
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
	return NgxConfD{NginxConf: string(file), ConfD: confFiles}
}
