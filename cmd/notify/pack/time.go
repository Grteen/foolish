package pack

import "time"

var Tz *time.Location
var TimeLayout string = "2006-01-02 15:04:05"

func InitTimeZone() {
	var err error
	Tz, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
}
