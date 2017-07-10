package httpUtil

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"time"

	"crypto/tls"
	"net/http/cookiejar"
)

//GetWebConFromUrl simply get web content
//from net
func GetWebConFromUrl(url string) (string, error) {
	response, err := DoRequest(url, nil, "GET", nil, 10*1000)
	if err != nil {
		return "", err
	}
	return getContentFromResponse(response)
}

//GetWebConFromUrlWithAllArgs get web content
//with some args
func GetWebConFromUrlWithAllArgs(url string, headerMap map[string]string, method string, postData []byte, timeOut time.Duration) (string, error) {
	response, err := DoRequest(url, headerMap, method, postData, timeOut)
	if err != nil {
		return "", err
	}
	return getContentFromResponse(response)
}

//GetWebConFromUrlWithHeader get web con from target url
//param headerMap is some header info
func GetWebConFromUrlWithHeader(url string, headerMap map[string]string) (string, error) {
	response, err := DoRequest(url, headerMap, "GET", nil, 10*1000)
	if err != nil {
		return "", err
	}
	return getContentFromResponse(response)
}

func getContentFromResponse(response *http.Response) (string, error) {
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

func DoRequest(url string, headerMap map[string]string, method string, postData []byte, timeOut time.Duration) (*http.Response, error) {
	timeout := time.Duration(timeOut * time.Millisecond)
	//https认证
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
	}

	client := http.Client{
		Timeout: timeout,
		Transport: tr,
	}
	client.Jar, _ = cookiejar.New(nil)
	method = strings.ToUpper(method)
	var req *http.Request
	var err error
	if postData != nil && method == "POST" {
		//print(string(postData))

		req, err = http.NewRequest(method, url, bytes.NewReader(postData))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
		if err != nil {
			return nil, err
		}
	} else {
		req, err = http.NewRequest(method, url, nil)
		if err != nil {
			return nil, err
		}

	}
	for key, value := range headerMap {
		req.Header.Add(key, value)
	}

	return client.Do(req)
}
