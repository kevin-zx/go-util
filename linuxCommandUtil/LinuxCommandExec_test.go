package linuxCommandUtil

import (
	"testing"
)

func TestLinuxServer_ExecCommand(t *testing.T) {
	ls := LinuxServer{UserName:"root",
		Password:"passwd",
		Host:"127.0.0.1",
		Port:22,
		}
	result,err := ls.ExecCommand("df -h")
	println("-----"+result+"-----")
	if err!=nil {
		panic(err)
	}

}