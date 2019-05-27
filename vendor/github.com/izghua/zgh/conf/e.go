package conf


var message = map[int]string{
	0 : "成功",
	100000001: "尚未设置错误码",
	400000000: "错误哦",
	400000001: "token失效,请检查后再试",

	500000000: "系统内部错误,请检查后再试",

}

func SetMsg(msg map[int]string) {
	for k,v := range msg {
		message[k] = v
	}
}

func GetMsg(code int) string {
	if msg, ok := message[code]; ok {
		return msg
	}
	return message[100000001]
}
