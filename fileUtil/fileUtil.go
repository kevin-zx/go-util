package fileUtil

import (
	"os"
	"bufio"
	//"fmt"
	"io/ioutil"
	"github.com/axgle/mahonia"
)



func CheckFileIsExist(fileName string) bool {
	var exist = true
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func CheckPathIsExist(pathName string) bool {
	var exist = true
	_,err := os.Stat(pathName)
	if err != nil && os.IsNotExist(err) {
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

//ioutil read file
func ReadFile(filename string, charset string) string{
	file,err := os.Open(filename)
	check(err)
	defer file.Close()
	decoder := mahonia.NewDecoder(charset)
	f := decoder.NewReader(file)
	file_content,err := ioutil.ReadAll(f)
	check(err)
	return string(file_content)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
