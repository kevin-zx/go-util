package dateUtil

import (
	"time"
	"github.com/kevin-zx/go-util/errorUtil"
)

//go的特殊format格式 1月2日 三点（二十四小时字） 四分 五秒 06年
var stand_date_format = "2006-01-02 15:04:05"

func GetCurrentDateTime() string {
	return normalFormatDate(time.Now())
}

func GetCurrentHour() int{
	return time.Now().Hour()
}

func normalFormatDate(date time.Time) string {
	return date.Format(stand_date_format)
}

//获取当前时间点增加或者减少一定时间后的字符串格式的日期
func GetDeltaDate(d time.Duration) string  {
	return normalFormatDate(time.Now().Add(d))
}

//从日期字符串格式转化成time.Time格式
func DateStr2Date(dateStr string) time.Time {
	loc,err := time.LoadLocation("Local")
	errorUtil.CheckErrorExit(err)
	theTime,err := time.ParseInLocation(stand_date_format,dateStr,loc)
	errorUtil.CheckErrorExit(err)
	return theTime
}