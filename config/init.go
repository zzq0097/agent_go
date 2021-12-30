package config

import (
	"agent/consts"
	"agent/dao"
	"agent/sql"
	"agent/util"
	"os"
	"strconv"
)

func Init() {
	CheckTable()
	ScanNgx()
	CheckParameter()
}

func CheckTable() {
	_, _ = util.DB.Exec(sql.CreateParameterTable)
	_, _ = util.DB.Exec(sql.CreateNgxLogTable)
	_, _ = util.DB.Exec(sql.CreateNgxStatusTable)
}

func ScanNgx() {
	files := map[string]string{
		consts.NgxDefCmdFile: consts.NgxDefConfFile,
		consts.NgxRpmCmdFile: consts.NgxRpmConfFile,
	}
	if value, err := dao.GetParameter(consts.BinName); err != nil || value == "" {
		for bin, conf := range files {
			if util.FileExist(bin) {
				dao.SaveOrUpdateParameter(consts.BinName, bin)
				dao.SaveOrUpdateParameter(consts.ConfName, conf)
				break
			}
		}
	}
}

func CheckParameter() {
	var size string
	if stat, err := os.Stat(consts.NgxRpmAccessLog); err == nil {
		size = strconv.FormatInt(stat.Size(), 10)
	} else {
		size = consts.MaxInt64Str
	}
	dao.SaveOrIgnoreParameter(consts.AccessLogOffset, size)
}
