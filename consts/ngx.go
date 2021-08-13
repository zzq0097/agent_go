package consts

const (
	Root   = "/home/nginxwise"
	DBFile = Root + "/db/sqlite.db"

	NgxDefRoot     = "/usr/local/nginx"
	NgxDefCmd      = NgxDefRoot + "/sbin"
	NgxDefCmdFile  = NgxDefCmd + "/nginx"
	NgxDefConf     = NgxDefRoot + "/conf"
	NgxDefConfFile = NgxDefConf + "/nginx.conf"

	NgxDefAccessLog = "/var/log/nginx/access.log"
	NgxDefErrorLog  = "/var/log/nginx/error.log"
)
