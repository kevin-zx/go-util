// 简单的文件下载的工具
// 不能超过100M
package simpleDownloadUtil

import (
	"net/url"
	"path"
	"os"
	"io"
	"net/http"
	"strings"
	"github.com/kevin-zx/go-util/fileUtil"
)


// 通过url下载文件
// fileUrl 是下载文件的地址
// savePath 是下载后存储的路径 未找到相关路径会return error
// saveName 是文件名 如果不存在则会新建 如果为空则会用默认的文件名
func SimpleDownLoadFile(fileUrl string,savePath string,saveName string) (error) {
	var err error
	// 对url进行格式化
	fileUrl = formatUrl(fileUrl)

	// 文件名为空通过url获取文件名
	if saveName == "" {
		saveName,err = GetFileNameFromUrl(fileUrl)
		if err != nil {
			return err
		}
	}

	// 检查路径是否存在
	if !fileUtil.CheckPathIsExist(savePath) {
		err := os.MkdirAll(savePath,os.ModePerm)
		if err != nil {
			return err
		}
	}

	// 下载
	res, err := http.Get(fileUrl)
	if err != nil {
		panic(err)
	}

	//存储
	if strings.HasSuffix(savePath,"/") || strings.HasSuffix(savePath,"\\") {
		savePath = savePath + string(os.PathSeparator)
	}
	f, err := os.Create(savePath+saveName)
	defer f.Close()
	defer res.Body.Close()
	if err != nil {
		return err
	}
	_,err = io.Copy(f, res.Body)


	return err
}

func formatUrl(fileUrl string) (string) {
	if strings.HasPrefix(fileUrl,"http://") || strings.HasPrefix(fileUrl,"https://") ||strings.HasPrefix(fileUrl,"ftp://") {
		return fileUrl
	}else if strings.HasPrefix(fileUrl, "//") {
		return "http:"+fileUrl
	}else{
		return "http://"+fileUrl
	}


}

//通过下载的url获取文件名
func GetFileNameFromUrl(fileRawURL string) (string,error){
	fileURL,err := url.Parse(fileRawURL)
	return path.Base(fileURL.Path),err
}


