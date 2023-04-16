package pack

func NotContainNil(target []interface{}) bool {
	for _, i := range target {
		if i == nil {
			return false
		}
	}
	return true
}
