package response

const (
	Success = 200
	Error   = 500

	RegisterEmailSendFailed = 10001 // 用户注册邮件发送失败
	PasswordIsNotIdentity   = 10002 // 密码不一致

	GetSmMsApiTokenFailed = 10003 // 获取sm.ms api token 失败
)
