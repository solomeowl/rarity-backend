package e

const (
	SUCCESS         = 200
	PARAMETER_ERROR = 400
	SERVER_ERROR    = 500
	NOT_FOUND       = 404
)

var MsgFlags = map[int]string{
	SUCCESS:         "ok",
	SERVER_ERROR:    "fail",
	PARAMETER_ERROR: "請求參數錯誤",
	NOT_FOUND:       "查無紀錄",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[SERVER_ERROR]
}
