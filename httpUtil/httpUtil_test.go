package httpUtil

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"
)

func testGetWebConFromUrl(t testing.T) {
	content, err := GetWebConFromUrl("http://www.oubeisiman.com")
	if err != nil {
		t.Error(err)
	}
	if len(content) == 0 {
		t.Error("get empty web content")
	}

}

func TestGetWebConFromUrl(t *testing.T) {
	//m.daishibuxi.cn
	urls := []string{"http://www.oubeisiman.com"}
	for _, sUrl := range urls {
		w, err := GetWebConFromUrlWithHeader(sUrl, map[string]string{"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36"})
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Println(len(w))

	}
	//go func() {
	//	for {
	//		reponse, err := GetWebResponseFromUrl("http://www.baidu.com")
	//		if err != nil {
	//			//t.Error(err)
	//			fmt.Println(err.Error())
	//			continue
	//		}
	//		content, err := ReadContentFromResponse(reponse, "utf-8")
	//		if err != nil {
	//			//t.Error(err)
	//			fmt.Println(err.Error())
	//			continue
	//		}
	//		if len(content) == 0 {
	//			fmt.Println("get empty web content")
	//			//t.Error("get empty web content")
	//		}
	//	}
	//
	//}()
	//time.Sleep(1000 * time.Second)
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
	} else {
		println(decodeString)
	}
}

func TestSendRequestWithProxy(t *testing.T) {
	re, err := SendRequestWithProxy("https://www.baidu.com", map[string]string{}, "GET", nil, 10*time.Second, "socks5://127.0.0.1:2080")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(re.StatusCode)
}
