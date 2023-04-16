package other

import (
	"be/pkg/errno"
	"strconv"
)

func NotContainNil(target []interface{}) bool {
	for _, i := range target {
		if i == nil {
			return false
		}
	}
	return true
}

// 将获取得到的空接口转化为string  如果包含的空接口数组内不含空 就返回 true
func ChangeNullItfToString(target []interface{}) ([]string, bool, error) {
	if NotContainNil(target) {
		res := make([]string, 0)
		for _, x := range target {
			r, ok := x.(string)
			if !ok {
				return nil, false, errno.ServiceFault
			}
			res = append(res, r)
		}
		return res, true, nil
	} else {
		return nil, false, nil
	}
}

// 将string转化为 int
func ChangeStringToInt(target []string) ([]int, bool, error) {
	res := make([]int, 0)
	for _, x := range target {
		r, err := strconv.Atoi(x)
		if err != nil {
			return nil, false, err
		}
		res = append(res, r)
	}
	return res, true, nil
}
