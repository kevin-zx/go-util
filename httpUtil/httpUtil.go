package httpUtil

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"time"

	"crypto/tls"
	"net/http/cookiejar"

	"net/url"
	"bufio"
	"golang.org/x/net/html/charset"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"errors"
	"fmt"
)

//GetWebConFromUrl simply get web content
//from net
func GetWebConFromUrl(url string) (string, error) {
	response, err := doRequest(url, nil, "GET", nil, 10*1000,"")
	if err != nil {
		return "", err
	}
	return GetContentFromResponse(response)
}

// get http.Response from url
func GetWebResponseFromUrl(url string) (*http.Response,error)  {
	return doRequest(url, nil, "GET", nil, 10*1000,"")

}

func GetWebResponseFromUrlWithHeader(url string,headerMap map[string]string,) (*http.Response,error)  {
	return doRequest(url, headerMap, "GET", nil, 10*1000,"")

}

//GetWebConFromUrlWithAllArgs get web content
//with some args
func GetWebConFromUrlWithAllArgs(url string, headerMap map[string]string, method string, postData []byte, timeOut time.Duration) (string, error) {
	response, err := doRequest(url, headerMap, method, postData, timeOut,"")
	if err != nil {
		return "", err
	}
	return GetContentFromResponse(response)
}

//GetWebConFromUrlWithHeader get web con from target url
//param headerMap is some header info
func GetWebConFromUrlWithHeader(url string, headerMap map[string]string) (string, error) {
	response, err := doRequest(url, headerMap, "GET", nil, 10*1000,"")
	if err != nil {
		return "", err
	}
	return GetContentFromResponse(response)
}

func GetContentFromResponse(response *http.Response) (string, error) {
	defer response.Body.Close()
	var c []byte
	for {
		buf := make([]byte, 1024)
		n, err := response.Body.Read(buf)
		if n == 0 {
			break
		}
		if err != nil && err != io.EOF {
			return "", err
		}

		c = append(c, buf[0:n]...)
	}
	return string(c), nil
}

func SendRequest(targetUrl string, headerMap map[string]string, method string, postData []byte, timeOut time.Duration) (*http.Response, error){
	return doRequest(targetUrl,headerMap,method,postData,timeOut,"")
}

func SendRequestWithProxy(targetUrl string, headerMap map[string]string, method string, postData []byte, timeOut time.Duration,proxy string) (*http.Response, error){
	return doRequest(targetUrl,headerMap,method,postData,timeOut,proxy)
}

func doRequest(targetUrl string, headerMap map[string]string, method string, postData []byte, timeOut time.Duration,proxy string) (*http.Response, error) {

	timeout := time.Duration(timeOut * time.Millisecond)
	//urli := url.URL{}
	//urlproxy, _ := urli.Parse("https://127.0.0.1:9743")
	//https认证
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
		//Proxy:http.ProxyURL(urlproxy),

	}
	if proxy != ""{
		urli := url.URL{}
		urlProxy, _ := urli.Parse(proxy)
		tr.Proxy = http.ProxyURL(urlProxy)
	}
	client := http.Client{
		Timeout: timeout,
		Transport: tr,
	}

	client.Jar, _ = cookiejar.New(nil)
	method = strings.ToUpper(method)
	var req *http.Request
	var err error

	if postData != nil && (method == "POST" || method == "PUT") {
		//print(string(postData))

		req, err = http.NewRequest(method, targetUrl, bytes.NewReader(postData))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
		if err != nil {
			return nil, err
		}
	} else {
		req, err = http.NewRequest(method, targetUrl, nil)
		if err != nil {
			return nil, err
		}

	}
	for key, value := range headerMap {
		req.Header.Add(key, value)
	}

	return client.Do(req)
}

func URLEncode(keyword string) (string)  {
	return url.QueryEscape(keyword)
}

// 这里他娘的Peek是针对Reader的 针对不了io.Reader 所以 io.Reader其实是进行了位移的
func detectContentCharset(body io.Reader) (string,*bufio.Reader) {
	r := bufio.NewReader(body)
	if data, err := r.Peek(1024); err == nil {
		if _, name, _ := charset.DetermineEncoding(data, ""); name!="" {
			return name,r
		}
	}
	return "utf-8",r
}
func ReadContentFromResponse(response *http.Response) (string, error) {
	defer response.Body.Close()
	if  response.StatusCode >= 300 || response.StatusCode < 200   {
		return "",errors.New(fmt.Sprintf("状态码为: %d",response.StatusCode))
	}
	char,data := detectContentCharset(response.Body)
	if data == nil {
		return "",errors.New("数据为空")
	}

	dec := mahonia.NewDecoder(char)
	preRd := dec.NewReader(data)
	preBytes,err := ioutil.ReadAll(preRd)
	if err != nil {
		return "",err
	}
	return string(preBytes),err
}
