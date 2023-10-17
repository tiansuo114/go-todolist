package e

var msgFlags = map[int]string{
	Success:       "ok",
	Error:         "fail",
	InvalidParams: "请求参数错误",

	SetPasswordDigestError: "设置密码失败",

	UserNotExist:     "用户不存在",
	UserAlreadyExist: "用户已存在",
	UserCreateError:  "创建用户失败",
}

func GetMsg(code uint) string {
	msg, ok := msgFlags[int(code)]
	if ok {
		return msg
	}

	return msgFlags[Error]
}
