package _jwt

import (
	"errors"
	"time"

	"github.com/jeffcail/go-im/core"

	"github.com/dgrijalva/jwt-go"
	"github.com/jeffcail/go-im/config"
)

type UserClaims struct {
	ID         int64  `json:"user_id"`
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
	Email      string `json:"email" validate:"required,email"`
	ClientType int    `json:"client_type"` // 用户登陆设备 1.web 2.pc 3.app 目前只支持web
	jwt.StandardClaims
}

// GenerateToken 生成token
func GenerateToken(userId int64, name, avatar, email string, clientType int) string {
	signKey := config.Config.Jwt.SECRET
	expireTime := config.Config.Jwt.Expire
	claims := UserClaims{
		ID:         userId,
		Name:       name,
		Avatar:     avatar,
		Email:      email,
		ClientType: clientType,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,
			ExpiresAt: time.Now().Unix() - expireTime,
			Issuer:    signKey,
		},
	}
	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenObj.SignedString([]byte(signKey))
	if err != nil {
		core.ImLogger.Error("Token 颁发失败")
	}
	return token
}

// ParseToken 解析token
func ParseToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return config.Config.Jwt.SECRET, nil
	})
	if err != nil {
		if v, ok := err.(*jwt.ValidationError); ok {
			if v.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("is not even a token")
			} else if v.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token is expired")
			} else if v.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token is not valid")
			} else {
				return nil, errors.New("token invalid")
			}
		}
	}
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("token invalid")
}

// RefreshToken 刷新token
func RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return config.Config.Jwt.SECRET, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		return tokenObj.SignedString([]byte(config.Config.Jwt.SECRET))
	}
	return "", errors.New("token invalid")
}
