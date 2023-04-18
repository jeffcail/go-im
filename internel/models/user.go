package model

import (
	"time"
)

type ImUser struct {
	ID              int64     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Name            string    `gorm:"column:name;NOT NULL"`
	Email           string    `gorm:"column:email"`
	EmailVerifiedAt time.Time `gorm:"column:email_verified_at"`
	Password        string    `gorm:"column:password"`
	RememberToken   string    `gorm:"column:remember_token"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
	Avatar          string    `gorm:"column:avatar"`                // 头像
	OauthID         string    `gorm:"column:oauth_id"`              // 第三方id
	BoundOauth      int       `gorm:"column:bound_oauth;default:0"` // 1\github 2\gitee
	DeletedAt       time.Time `gorm:"column:deleted_at"`
	OauthType       int       `gorm:"column:oauth_type"`       // 1.微博 2.github
	Status          int       `gorm:"column:status;default:0"` // 0 离线 1 在线
	Bio             string    `gorm:"column:bio"`              // 用户简介
	Sex             int       `gorm:"column:sex;default:0"`    // 0 未知 1.男 2.女
	ClientType      int       `gorm:"column:client_type"`      // 1.web 2.pc 3.app
	Age             int       `gorm:"column:age"`
	LastLoginTime   time.Time `gorm:"column:last_login_time"` // 最后登录时间
}

func (u *ImUser) GetTable() string {
	return "im_user"
}
