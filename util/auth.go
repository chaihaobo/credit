// Package util
// @author： Boice
// @createTime：2022/11/29 16:27
package util

import (
	"credit-platform/constant"
	"crypto/md5"
	"encoding/hex"
)

//	AuthCheckPassword 验证密码和签名后的密码是否一致
func AuthCheckPassword(inputPassword string, passwordHash string) bool {
	hash := md5.New()
	hash.Write([]byte(inputPassword + constant.SaltPassword))
	inputPasswordSum := hex.EncodeToString(hash.Sum(nil))
	return inputPasswordSum == passwordHash
}
