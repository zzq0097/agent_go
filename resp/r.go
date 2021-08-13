package resp

type R struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const OK = 0
const ERROR = 1

func Success() R {
	return R{OK, "success", nil}
}

func SuccessData(data interface{}) R {
	return R{OK, "success", data}
}

func Error() R {
	return R{ERROR, "error", nil}
}

func ErrorMsg(msg string) R {
	return R{ERROR, msg, nil}
}
