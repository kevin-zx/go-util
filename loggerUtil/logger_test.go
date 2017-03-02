package loggerUtil

import (
	"testing"
	"github.com/kevin-zx/go-util/fileUtil"
)

func TestLogUtil(t *testing.T) {
	var logFilePath string = "test.log"
	SetConsole(true)
	SetRollingFile("D:/", logFilePath, 10, 5, KB)
	SetLevel(DEBUG)
	Info("test")

	if fileUtil.CheckFileIsExist("D:/test.log") {
		println("success")
	}else {
		//println("success")
		t.Fail()
		t.Error("no such file")
	}

}
