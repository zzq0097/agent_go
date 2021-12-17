package consts

const (
	Root   = "/opt/agent"
	DBFile = Root + "/db/sqlite.db"

	NgxDefRoot      = "/usr/local/nginx"
	NgxDefCmd       = NgxDefRoot + "/sbin"
	NgxDefCmdFile   = NgxDefCmd + "/nginx"
	NgxDefConf      = NgxDefRoot + "/conf"
	NgxDefConfFile  = NgxDefConf + "/nginx.conf"
	NgxDefAccessLog = NgxDefRoot + "/logs/access.log"
	NgxDefErrorLog  = NgxDefRoot + "/logs/error.log"

	NgxRpmRoot      = "/etc/nginx"
	NgxRpmCmdFile   = "/usr/sbin/nginx"
	NgxRpmConf      = NgxRpmRoot
	NgxRpmConfFile  = NgxDefConf + "/nginx.conf"
	NgxRpmAccessLog = "/var/log/nginx/access.log"
	NgxRpmErrorLog  = "/var/log/nginx/error.log"
)
