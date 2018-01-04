package server_monitor

import (
	"github.com/kevin-zx/go-util/linuxCommandUtil"
	"testing"
	//"golang.org/x/crypto/ssh"
	//"strings"
	//"fmt"
)

func TestLinuxMonitor_MonitorDisk(t *testing.T) {
	//UserName string
	//Password string
	//PrivateKeyFilePath string
	//Host     string
	//Port     int
	//client   ssh.Client
	lm := LinuxMonitor{LinuxServer: linuxCommandUtil.LinuxServer{UserName: "root",
		Password: "Iknowthat@@!221",
		Host:     "115.159.79.85",
		Port:     22}}
	lm.MonitorDisk()

}
