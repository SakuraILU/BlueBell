package control

type Code int

const (
	InvalidParam Code = iota
	UserNotExist
	UserExist
	InvalidPassword
	NotLogin
	InvalidToken

	Success
)

var error2text = map[Code]string{
	InvalidParam:    "请求参数有误",
	UserNotExist:    "用户不存在",
	UserExist:       "用户已存在",
	InvalidPassword: "密码错误",
	NotLogin:        "未登录",
	InvalidToken:    "非法校验码",

	Success: "请求成功",
}

func (c Code) string() string {
	return error2text[c]
}
