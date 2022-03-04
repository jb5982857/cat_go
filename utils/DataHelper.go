package dataHelper

import "time"

var cstSh, _ = time.LoadLocation("Asia/Shanghai")

func GetCurrentDataTimeToDb() time.Time {
	return time.Now().In(cstSh)
}
