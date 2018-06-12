package dateUtil

import (
	"testing"
	"time"
)

func TestGetDeltaDate(t *testing.T) {
	println(GetDeltaDateTime(-10 * time.Hour))
}

func TestDateStr2Date(t *testing.T) {
	//println(time.Now().Unix())
	mtime,err := DateStr2Date("2015-11-03 03:04:05")
	if err!=nil {
		t.Error(err)
	}else{
		println(mtime.Unix())
	}
}