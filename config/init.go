package config

import (
	"agent/consts"
	"agent/dao"
	"agent/util"
	"fmt"
)

func Init() {
	dao.CreateTable()
	ScanNgx()
}

func ScanNgx() {
	files := map[string]string{
		consts.NgxDefCmdFile: consts.NgxDefConfFile,
		consts.NgxRpmCmdFile: consts.NgxRpmConfFile,
	}
	if value, err := dao.GetParameter(consts.BinName); err != nil || value == "" {
		for bin, conf := range files {
			fmt.Println(bin)
			if util.FileExist(bin) {
				dao.SetParameter(consts.BinName, bin)
				dao.SetParameter(consts.ConfName, conf)
				break
			}
		}
	}
}
