package linuxCommandUtil

import (
	"testing"
)

func TestLinuxServer_ExecCommand(t *testing.T) {
	ls := LinuxServer{UserName:"root",
		Password:"Iknowthat@@!221",
		Host:"115.159.79.85",
		Port:22,
		}
	result,err := ls.ExecCommand("df -h")
	println("-----"+result+"-----")
	if err!=nil {
		panic(err)
	}

}