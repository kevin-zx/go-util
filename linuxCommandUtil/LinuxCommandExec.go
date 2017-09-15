package linuxCommandUtil

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"bytes"

	"io/ioutil"
)

type LinuxServer struct {
	//login linux server username e.g. root
	UserName string
	//linux server password
	Password string
	// if use privateKey to auth login, than set PrivateKeyFilePath
	PrivateKeyFilePath string
	// linux server ip address e.g. 192.168.0.112
	Host     string
	// ssh port normal port is 22 e.g. 22
	Port     int
	// ssh client to act new session ...
	client   ssh.Client
}

//login linux server
func (ls *LinuxServer) login() error {
	authMethod := []ssh.AuthMethod{}
	//if use privateKey to login
	if ls.PrivateKeyFilePath != ""{
		authMethod = []ssh.AuthMethod{PublicKeyFile(ls.PrivateKeyFilePath)}
	}else{
		// use password login
		authMethod = []ssh.AuthMethod{ssh.Password(ls.Password)}
	}

	conf := ssh.ClientConfig{User: ls.UserName,
		Auth:                      authMethod,
		// there must be this line
		HostKeyCallback:           ssh.InsecureIgnoreHostKey(),
		//HostKeyCallback: ssh.FixedHostKey()
		}
	//handshake with server
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", ls.Host, ls.Port), &conf)
	if err != nil {
		return err
	}
	ls.client = *client
	return nil
}

// loader public key by file
func PublicKeyFile(file string) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}
	return ssh.PublicKeys(key)
}

func (ls *LinuxServer) ExecCommandInOneSession() error {
	//@todo this function can exec command one by one in one session so you can exec a series command
	return nil
}

func (ls *LinuxServer) CloseSession() error {
	//@todo this funtion is close for ExecCommandInOneSession()
	return nil
}

// exec command and return exec result
func (ls *LinuxServer) ExecCommand(command string) (string, error) {
	// if conn not establish, than establish the conn
	if ls.client.Conn == nil {
		err := ls.login()
		if err != nil {
			return "",err
		}
	}

	defer ls.Close()

	//make a new session to exec command
	session, err := ls.client.NewSession()
	if err != nil {
		return "",err
	}
	defer session.Close()
	b := bytes.NewBuffer(make([]byte, 0))
	//put exec result to b buffer
	session.Stdout = b
	err = session.Run(command)
	return b.String(),err
}

func (ls *LinuxServer) Close() error {
	if ls != nil {
		err := ls.client.Close()
		return err
	}
	return nil
}
