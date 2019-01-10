package fileUtil

import (
	"testing"
	"log"
)

func TestCheckPathIsExist(t *testing.T) {
	testMap := map[string]bool{
		"d:":true,
		"d:/adfkjaslkdfjkladjf":false,
	}
	for path,expect := range testMap{
		if CheckPathIsExist(path) != expect {
			log.Printf("path:%s 期望的存在结果是 %b 但是程序判断的存在状态是 %b",path,expect,!expect)
			t.Fail()
		}
	}
}
