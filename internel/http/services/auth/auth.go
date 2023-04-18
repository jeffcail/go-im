package auth

import (
	"errors"

	"github.com/jeffcail/go-im/core"
	model "github.com/jeffcail/go-im/internel/models"
)

// FindOneUserByName 获取用户信息
func FindOneUserByName(name string) (*model.ImUser, error) {
	u := &model.ImUser{}
	err := core.Db.Where("name = ?", name).Find(u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

// RegisteredService 用户注册
func RegisteredService(email, name, password, avatar string) error {
	u := &model.ImUser{
		Name:     name,
		Email:    email,
		Password: password,
		Avatar:   avatar,
	}
	err := core.Db.Create(u).Error
	if err != nil {
		return errors.New("注册账号失败")
	}
	return nil
}

// UpdateUserInfo 修改用户信息
func UpdateUserInfo(user *model.ImUser) error {
	return core.Db.Model(&model.ImUser{}).Where("name = ?", user.Name).Save(user).Error
}
