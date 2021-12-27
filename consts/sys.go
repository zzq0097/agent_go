package consts

const (
	Root   = "/opt/agent"
	DBFile = Root + "/db/sqlite.db"

	BinName  = "bin_file"
	ConfName = "conf_file"

	NgxDefRoot      = "/usr/local/nginx/"
	NgxDefCmd       = NgxDefRoot + "sbin/"
	NgxDefCmdFile   = NgxDefCmd + "nginx"
	NgxDefConf      = NgxDefRoot + "conf/"
	NgxDefConfD     = NgxDefConf + "conf.d/"
	NgxDefConfFile  = NgxDefConf + "nginx.conf"
	NgxDefLog       = NgxDefRoot + "logs/"
	NgxDefAccessLog = NgxDefRoot + "access.log"
	NgxDefErrorLog  = NgxDefRoot + "error.log"

	NgxRpmRoot      = "/etc/nginx/"
	NgxRpmCmdFile   = "/usr/sbin/nginx"
	NgxRpmConf      = NgxRpmRoot
	NgxRpmConfD     = NgxRpmConf + "conf.d/"
	NgxRpmConfFile  = NgxRpmConf + "nginx.conf"
	NgxRpmLog       = "/var/log/nginx/"
	NgxRpmAccessLog = NgxRpmLog + "access.log"
	NgxRpmErrorLog  = NgxRpmLog + "error.log"
)
