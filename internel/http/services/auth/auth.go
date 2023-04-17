package auth

import (
	"errors"

	"github.com/jeffcail/go-im/core"
	model "github.com/jeffcail/go-im/internel/models"
)

// FindOneUserByName
func FindOneUserByName(name string) (*model.ImUser, error) {
	u := &model.ImUser{}
	err := core.Db.Where("name = ?", name).Find(u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

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
