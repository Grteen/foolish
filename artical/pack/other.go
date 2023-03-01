package pack

import "be/pkg/errno"

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
