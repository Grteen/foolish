package check

import (
	"regexp"
)

// 仅验证数据合法性 不验证一致性

// 检测用户名称 如果合法返回 true
func CheckUserName(username string) bool {
	if len(username) < 3 || len(username) > 18 {
		return false
	}
	userPwReg := regexp.MustCompile("[^0-9a-zA-Z\\-_]")
	return !userPwReg.MatchString(username)
}

// 检测用户密码 （未加密） 如果合法返回 true
func CheckUserPassWord(password string) bool {
	if len(password) < 3 || len(password) > 18 {
		return false
	}
	userPwReg := regexp.MustCompile("[^0-9a-zA-Z\\-_]")
	return !userPwReg.MatchString(password)
}

// 检测用户邮箱 如果合法返回true
func CheckUserEmail(email string) bool {
	emailReg := regexp.MustCompile("[^0-9a-zA-Z@.]")
	if emailReg.MatchString(email) {
		return false
	}
	emailReg = regexp.MustCompile(".+(@qq.com)$")
	return emailReg.MatchString(email)
}

// 检测通知内容 如果合法返回true
func CheckNotifyText(text string) bool {
	if len(text) > 1000 {
		return false
	}

	reg := regexp.MustCompile("[^\u4e00-\u9fa5a-z0-9A-Z_ ?,.;:\\[\\]\\{\\}\\-]")
	return !reg.MatchString(text)
}

// 检测输入是否为正数 如果是返回 true
func CheckPostiveNumber(val int32) bool {
	return val > 0
}

func CheckZeroOrPostive(val int32) bool {
	return val >= 0
}

// 检测输入的string长度是否合法 如果是返回 true
func CheckStringLength(val string) bool {
	return len(val) != 0
}

// 检测输入的string数组中所有string长度是否合法 如果是返回 true
func CheckStringArray(val []string) bool {
	for _, str := range val {
		if len(str) == 0 {
			return false
		}
	}
	return true
}

// 检测输入的数组是否全为正数 如果是返回 true
func CheckPostiveArray(val []int32) bool {
	if len(val) == 0 {
		return false
	}
	for _, v := range val {
		if v < 0 {
			return false
		}
	}
	return true
}

// 检测动态文本是否合法 如果是返回true
func CheckActionText(text string) bool {
	if len(text) <= 0 || len(text) > 1000 {
		return false
	}
	return true
}

// 检测评论长度是否合法 如果是返回true
func CheckCommentText(comment string) bool {
	return len(comment) <= 500
}

// 检测收藏夹权限是否合法 如果是返回true
func CheckStarFolderPublic(public int32) bool {
	var a struct{}
	mp := map[int32]struct{}{0: a, 1: a}
	_, ok := mp[public]
	return ok
}

// 检测 粉丝关注公开权限 如果合法 返回 true
func CheckPublic(public int32) bool {
	var a struct{}
	mp := map[int32]struct{}{0: a, 1: a}
	_, ok := mp[public]
	return ok
}

// 检测 description 如果合法 返回 true
func CheckDesc(desc string) bool {
	if len(desc) > 500 {
		return false
	}
	descReg := regexp.MustCompile("[^\u4e00-\u9fa5a-z0-9A-Z_ ?,.;:\\[\\]\\{\\}\\-]")

	return !descReg.MatchString(desc)
}

// 检测 search 中的 keyword 如果合法 返回 true
func CheckKeyWord(key string) bool {
	return len(key) <= 30
}

// 检测search中的limit 如果合法 返回 true
func CheckSearchLimit(limit int32) bool {
	if limit < 0 || limit > 20 {
		return false
	}
	return true
}

// 检测offset 如果合法 返回 true
func CheckOffset(offset int32) bool {
	return offset >= 0
}

// 检测公告参数是否合法
func CheckPubNoticeText(text string) bool {
	return len(text) <= 500
}

// 检测公告参数 如果合法 返回true
func CheckPubNotice(username, text string) bool {
	if !CheckUserName(username) || !CheckPubNoticeText(text) {
		return false
	}
	return true
}
