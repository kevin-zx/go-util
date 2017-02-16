package fileUtil

import (
	"os"
	"bufio"
	"fmt"
)

func CheckFileIsExist(file_name string) bool {
	var exsit = true
	if _, err := os.Stat(file_name); os.IsNotExist(err) {
		exsit = false
	}
	return exsit
}

func WriteToFile(file_name string, is_append bool, content string) {
	var f *os.File
	var err error
	if CheckFileIsExist(file_name) {
		f, err = os.OpenFile(file_name, os.O_CREATE, 0666)
		defer f.Close()
	} else {
		f, err = os.Create(file_name)
	}
	check(err)
	w := bufio.NewWriter(f)  //创建新的 Writer 对象
	n4, err:= w.WriteString("bufferedn")
	fmt.Printf("写入 %d 个字节n", n4)
	w.Flush()
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
