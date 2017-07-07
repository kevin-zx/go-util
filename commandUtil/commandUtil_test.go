package commandUtil

import "testing"

func TestExecCommand(t *testing.T) {
	ExecCommand("ping",[]string{"192.168.0.1"})

}