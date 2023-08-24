package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// GeneratePassword 加密明文密码
func GeneratePassword(pwd string) (string, error) {
	/*
		自适应哈希算法是一种根据数据类型和大小自动选择哈希算法的技术，以提高哈希效率和安全性。
		下面是使用自适应哈希算法的一般步骤：
			选择一个自适应哈希算法库，并在代码中引入该库。
			创建一个哈希对象。将要哈希的数据传递给哈希对象的输入流。
			调用哈希对象的哈希计算方法，该方法将返回哈希值。
	*/
	encryptedPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(encryptedPwd), err
}

// ValidatePassword 验证密码
// pwd 明文密码
// encryptedPwd 加密的密码
func ValidatePassword(pwd string, encryptedPwd string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(encryptedPwd), []byte(pwd)); err != nil {
		return false, errors.New("wrong password")
	}
	return true, nil
}
