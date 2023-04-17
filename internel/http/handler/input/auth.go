package input

type SendRegisterEmailInput struct {
	Email string `json:"email" validate:"required,email"`
}

// RegisteredInput 注册接收参数
type RegisteredInput struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required"`
	Code            string `json:"code" validate:"required"`
	Password        string `json:"password" validate:"required"`
	PasswordConfirm string `json:"password_confirm" validate:"required"`
}
