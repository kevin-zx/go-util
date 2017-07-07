package commandUtil

import (
	"os/exec"
	"fmt"
	"bufio"
	"io"
)

func KillProcess(processName string) {
	ExecCommand("taskkill", []string{"/IM", processName, "/F"})
}

func ExecCommand(commandName string, params []string) (string, bool) {
	cmd := exec.Command(commandName, params...)
	fmt.Println(cmd.Args)
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		return "", false
	}

	err = cmd.Start()
	if err != nil {
		// panic(err)
		fmt.Println(err)
	}
	reader := bufio.NewReader(stdout)
	//实时循环读取输出流中的一行内容
	commandResult := ""
	for {
		line, err2 := reader.ReadString('\n')

		if err2 != nil || io.EOF == err2 {
			break
		}
		commandResult += line + "\r\n"
		print(line)
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println(err)
	}
	return commandResult, true
}
func ExecWithoutWait(commandName string, params []string) {
	// fmt.Println("1")
	cmd := exec.Command(commandName, params...)
	fmt.Println(cmd.Args)
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println("2")
}
