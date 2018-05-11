package httpUtil

import (
	"encoding/json"
	"testing"
	"time"
	"log"
)

func testGetWebConFromUrl(t testing.T) {
	content, err := GetWebConFromUrl("https://www.baidu.com")
	if err != nil {
		t.Error(err)
	}
	if len(content) == 0 {
		t.Error("get empty web content")
	}

}

func testGetWebConFromUrlWithHeader(t testing.T) {
	header := make(map[string]string)
	header["Accept-Encoding"] = "gzip, deflate, sdch"
	header["Accept-Language"] = "en-US,en;q=0.8"
	header["User-Agent"] = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.116 Safari/537.36"
	header["Accept"] = "*/*"
	header["Referer"] = "http://www.jd.com/"
	header["Connection"] = "keep-alive"
	url := "http://dd-search.jd.com/?ver=2&zip=1&key=s&pvid=4a6j59vi.tnm8dd&t=1478588720776&curr_url=www.jd.com%2F&callback=jQuery1880422"
	content, err := GetWebConFromUrlWithHeader(url, header)
	if err != nil {
		t.Error(err)
	}
	if len(content) == 0 {
		t.Error("get empty web content")
	}
}

func testGetWebConFromUrlWithAllArgs(t testing.T) {
	header := make(map[string]string)
	header["Accept-Encoding"] = "gzip, deflate, sdch"
	header["Accept-Language"] = "en-US,en;q=0.8"
	header["User-Agent"] = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.116 Safari/537.36"
	header["Accept"] = "*/*"
	header["Referer"] = "http://www.jd.com/"
	header["Connection"] = "keep-alive"
	url := "http://dd-search.jd.com/?ver=2&zip=1&key=s&pvid=4a6j59vi.tnm8dd&t=1478588720776&curr_url=www.jd.com%2F&callback=jQuery1880422"
	method := "get"
	var pastStruct = []struct {
		a string
		b string
	}{
		{"2", "3"},
		{"3", "4"},
	}
	postData, err := json.Marshal(pastStruct)
	if err != nil {
		t.Error(err.Error())
	}
	timeOut := time.Duration(10)
	content, err := GetWebConFromUrlWithAllArgs(url, header, method, postData, timeOut)
	if err != nil {
		t.Error(err.Error())
	}
	if len(content) == 0 {
		t.Error("get empty web content")
	}


}

func TestURLEncode(t *testing.T) {
	decodeString := URLEncode("安瓶")
	if decodeString != "%E5%AE%89%E7%93%B6" {
		log.Println(decodeString)
		t.Fail()
	}else{
		println(decodeString)
	}
}