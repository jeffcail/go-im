package utils

import (
	rand2 "crypto/rand"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// RandomEmailCode 生成注册邮箱验证码
func RandomEmailCode() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}

func HashEncrypt(str string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func Random(m int64) int {
	max := big.NewInt(m)
	i, err := rand2.Int(rand2.Reader, max)
	if err != nil {
		log.Fatal("rand: ", err)
	}
	return i.BitLen()
}
