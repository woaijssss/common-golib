package common

const (
	SUCCESS       = 200 //成功
	ERROR         = 500 //失败
	InvalidParams = 400 //参数错误
)

var MsgFlags = map[int]string{
	SUCCESS:       "成功",
	ERROR:         "失败",
	InvalidParams: "参数不合法",
}

// GetMsg get error information based on Code
//func GetMsg(code int) string {
//	msg, ok := MsgFlags[code]
//	if ok {
//		return msg
//	}
//
//	return MsgFlags[ERROR]
//}
