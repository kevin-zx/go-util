package fileUtil

import (
	"os"
	"bufio"
	//"fmt"
)

func CheckFileIsExist(file_name string) bool {
	var exist = true
	if _, err := os.Stat(file_name); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

//写入数据到文件返回写入文件得字长
func WriteToFile(file_name string, is_append bool, content string) int {
	var f *os.File
	var err error
	if CheckFileIsExist(file_name) {
		if is_append {
			f, err = os.OpenFile(file_name, os.O_APPEND, 0666)
		}else {
			f, err = os.OpenFile(file_name, os.O_CREATE, 0666)
		}
		defer f.Close()
	} else {
		f, err = os.Create(file_name)
	}
	check(err)
	w := bufio.NewWriter(f)  //创建新的 Writer 对象
	n4, err:= w.WriteString(content)
	w.Flush()
	check(err)
	return n4
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
