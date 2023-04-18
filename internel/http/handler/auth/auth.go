package auth

import (
	"fmt"
	"time"

	"go.uber.org/zap"

	model "github.com/jeffcail/go-im/internel/models"

	_jwt "github.com/jeffcail/go-im/pkg/jwt"

	"github.com/spf13/cast"

	"github.com/jeffcail/go-im/internel/http/services/auth"

	"github.com/jeffcail/wildrocket/email"

	"github.com/jeffcail/go-im/core"

	"github.com/jeffcail/go-im/config"

	"github.com/jeffcail/go-im/utils"

	"github.com/gin-gonic/gin"
	"github.com/jeffcail/go-im/internel/http/handler/inputOut"
	"github.com/jeffcail/go-im/pkg/response"
	"github.com/jeffcail/wildrocket/vali"
)

type (
	AuthHandler struct{}
)

// SendRegisterEmail 发送邮件验证
func (a *AuthHandler) SendRegisterEmail(c *gin.Context) {
	input := &inputOut.SendRegisterEmailInput{}
	_ = c.BindJSON(input)
	msg := vali.BindValidate(input)
	if msg != "" {
		response.ToJsonResponse(response.Error, msg).ToJson(c)
		return
	}
	go sendEmail(input.Email)

	response.ToJsonResponse(response.Success, "验证码发送成功").ToJson(c)
	return
}

func sendEmail(em string) {
	emailCode := utils.RandomEmailCode()
	title := fmt.Sprintf("邮箱为:%s注册IM邮件", em)
	content := emailTemplate(emailCode)
	if err := email.SendMail(
		config.Config.Email.MailFrom,
		em,
		title,
		content,
		config.Config.Email.MailServer,
		config.Config.Email.MailServerPort,
		config.Config.Email.MailPassword,
	); err != nil {
		core.ImLogger.Error(fmt.Sprintf("email: %s 发送验证码失败 err: %v", em, err))
	}
	if err := core.RDB.Set(em, emailCode, time.Minute*5).Err(); err != nil {
		core.ImLogger.Error(fmt.Sprintf("email: %s 发送验证码失败 err: %v", em, err))
	}
}

// Registered 注册
func (a *AuthHandler) Registered(c *gin.Context) {
	input := &inputOut.RegisteredInput{}
	_ = c.Bind(input)
	msg := vali.BindValidate(input)
	if msg != "" {
		response.ToJsonResponse(response.Error, msg).ToJson(c)
		return
	}
	if input.Password != input.PasswordConfirm {
		response.ToJsonResponse(response.PasswordIsNotIdentity, "两次密码不一致").ToJson(c)
		return
	}

	user, err := auth.FindOneUserByName(input.Name)
	if err != nil {
		response.ToJsonResponse(response.Error, cast.ToString(err)).ToJson(c)
		return
	}
	if user.ID >= 0 {
		response.ToJsonResponse(response.Error, "用户名已经存在").ToJson(c)
		return
	}

	p, err := utils.HashEncrypt(input.Password)
	if err != nil {
		core.ImLogger.Error(fmt.Sprintf("%s:%s密码加密失败, 密码为: %s", input.Name, input.Email, input.Password))
	}

	err = auth.RegisteredService(input.Email, input.Name, p, avatar[utils.Random(6)])
	if err != nil {
		response.ToJsonResponse(response.PasswordIsNotIdentity, cast.ToString(err)).ToJson(c)
		return
	}

	response.ToJsonResponse(response.Success, "注册成功").ToJson(c)
	return
}

// Login
func (a *AuthHandler) Login(c *gin.Context) {
	input := &inputOut.LoginInput{}
	_ = c.BindJSON(input)
	msg := vali.BindValidate(input)
	if msg != "" {
		response.ToJsonResponse(response.Error, msg).ToJson(c)
		return
	}
	u, err := auth.FindOneUserByName(input.Name)
	if err != nil {
		response.ToJsonResponse(response.Error, "登陆异常，请稍后再试！").ToJson(c)
		return
	}
	if u.ID == 0 {
		response.ToJsonResponse(response.Error, "您还未注册，请先注册！").ToJson(c)
		return
	}
	if ok := utils.ComparePassword(u.Password, input.Password); !ok {
		response.ToJsonResponse(response.Error, "密码不正确，请输入正确的密码！").ToJson(c)
		return
	}

	// 下发token
	token := _jwt.GenerateToken(u.ID, u.Name, u.Avatar, u.Email, input.ClientType)
	data := userReturnInfo(token, u)
	// 单点登陆，将别的登陆状态挤下线

	// 修改用户最后登陆时间
	u.LastLoginTime = time.Now()
	if err = auth.UpdateUserInfo(u); err != nil {
		core.ImLogger.Error(fmt.Sprintf("用户:%s登陆修改最后登陆时间失败", u.Name), zap.Error(err))
	}

	response.ToJsonResponse(response.Success, "登陆成功", data).ToJson(c)
	return
}

func userReturnInfo(token string, user *model.ImUser) *inputOut.UserReturnInformation {
	return &inputOut.UserReturnInformation{
		ID:         user.ID,
		Name:       user.Name,
		Avatar:     user.Avatar,
		Email:      user.Email,
		Token:      token,
		ExpireTime: config.Config.Jwt.Expire,
		Sex:        user.Sex,
		ClientType: user.ClientType,
		Bio:        user.Bio,
	}
}
