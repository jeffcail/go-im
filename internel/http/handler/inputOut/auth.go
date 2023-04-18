package inputOut

type SendRegisterEmailInput struct {
	Email string `json:"email" validate:"required,email"`
}

// RegisteredInput 注册接收参数
type RegisteredInput struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Code            string `json:"code" validate:"required"`
	Password        string `json:"password" validate:"required"`
	PasswordConfirm string `json:"password_confirm" validate:"required"`
}

// LoginInput 登陆
type LoginInput struct {
	Name       string `json:"name" validate:"required"`
	Password   string `json:"password" validate:"required"`
	ClientType int    `json:"client_type" validate:"required"`
}

type UserReturnInformation struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
	Email      string `json:"email"`
	Token      string `json:"token"`
	ExpireTime int64  `json:"expire_time"`
	Sex        int    `json:"sex"`
	ClientType int    `json:"client_type"`
	Bio        string `json:"bio"`
}
