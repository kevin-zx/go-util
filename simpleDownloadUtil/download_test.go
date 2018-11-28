package simpleDownloadUtil

import (
	"testing"
	"log"
)

func TestSimpleDownLoadFile(t *testing.T) {
	err := SimpleDownLoadFile("//img2.aichuangfu.cn/upload/vod/2014-04-21/13980649420.jpg","./","")
	if err != nil {
		t.Fatal(err)
	}

}

func TestGetFileNameFromUrl(t *testing.T) {
	testMap := map[string]string{
		"//img2.aichuangfu.cn/upload/vod/2014-04-21/13980649420.jpg":"13980649420.jpg",
		"http://techforum-img.cn-hangzhou.oss-pub.aliyun-inc.com/1528185610558/JAVA1.4.pdf?Expires=1533701900&OSSAccessKeyId=LTAIAJ2WBIdlRPlb&Signature=b8is%2BUYggn0sHA6LHNqtr3QciAs%3D":"JAVA1.4.pdf",
	}
	for fileRawURL,expectFileName := range testMap {
		fileName,err := GetFileNameFromUrl(fileRawURL)
		if err != nil {
			t.Fatal(err)
		}
		if fileName != expectFileName {
			log.Printf("出现错误：期望的文件名是 %s, 得到的文件名是 %s",expectFileName,fileName,)
			t.Fail()
		}

	}
}