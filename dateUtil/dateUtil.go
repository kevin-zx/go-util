package dateUtil

import "time"

//go的特殊format格式 1月2日 三点（二十四小时字） 四分 五秒 06年
var stand_date_format = "2006-01-02 15:04:05"

func GetCurrentDateTime() string {
	return time.Now().Format(stand_date_format)
}