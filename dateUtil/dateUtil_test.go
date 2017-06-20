package dateUtil

import (
	"testing"
	"time"
)

func TestGetDeltaDate(t *testing.T) {
	println(GetDeltaDate(-10 * time.Hour))
}

func TestDateStr2Date(t *testing.T) {
	//println(time.Now().Unix())
	println(DateStr2Date("2015-11-03 03:04:05").Unix())
}