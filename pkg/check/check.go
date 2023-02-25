package check

import (
	"regexp"
)

// 仅验证数据合法性 不验证一致性

func CheckUserName(username string) bool {
	if len(username) < 3 || len(username) > 18 {
		return false
	}
	userPwReg := regexp.MustCompile("[^0-9a-zA-Z\\-_]")
	if userPwReg.MatchString(username) {
		return false
	}
	return true
}

func CheckUserPassWord(password string) bool {
	if len(password) < 3 || len(password) > 18 {
		return false
	}
	userPwReg := regexp.MustCompile("[^0-9a-zA-Z\\-_]")
	if userPwReg.MatchString(password) {
		return false
	}
	return true
}

func CheckUserEmail(email string) bool {
	emailReg := regexp.MustCompile("[^0-9a-zA-Z@.]")
	if emailReg.MatchString(email) {
		return false
	}
	emailReg = regexp.MustCompile(".+(@qq.com)$")
	if !emailReg.MatchString(email) {
		return false
	}
	return true
}

func CheckPostiveNumber(val int32) bool {
	if val < 0 {
		return false
	}
	return true
}
