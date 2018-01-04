package dateUtil

import (
	"time"
)

//go的特殊format格式 1月2日 三点（二十四小时字） 四分 五秒 06年
var stand_date_time_format = "2006-01-02 15:04:05"

var stand_date_format = "2006-01-02"

func GetCurrentDateTime() string {
	return normalFormatDateTime(time.Now())
}

func GetCurrentHour() int{
	return time.Now().Hour()
}

func normalFormatDateTime(date time.Time) string {
	return date.Format(stand_date_time_format)
}
func normalFormatDate(date time.Time) string {
	return date.Format(stand_date_format)
}
//获取当前时间点增加或者减少一定时间后的字符串格式的日期
func GetDeltaDateTime(d time.Duration) string  {
	return normalFormatDateTime(time.Now().Add(d))
}

func GetDeltaDate(d time.Duration) string  {
	return normalFormatDate(time.Now().Add(d))
}

//从日期字符串格式转化成time.Time格式
func DateStr2Date(dateStr string) (targetTime time.Time,err error) {
	loc,err := time.LoadLocation("Local")
	return time.Time{},err
	targetTime,err = time.ParseInLocation(stand_date_time_format,dateStr,loc)
	return

}