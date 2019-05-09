package wd_crawler

import (
	"fmt"
	"github.com/kevin-zx/go-util/mysqlUtil"
	"testing"
	"time"
)

func TestWdRequest_GetContent(t *testing.T) {
	header := make(map[string]string)
	mu := mysqlutil.MysqlUtil{}
	mu.InitMySqlUtilDetail("182.254.155.218", 3306, "spider_center", "spiderdb@wd", "spider", 1, 1)
	header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.139 Safari/537.36"
	wr := WdRequest{Header: header, mu: mu, Redirect: 1, Single: 0, ConcurrentNum: 0, StoreType: 1, ContentType: 1}
	wr.SendRequest("http://www.xiaohongshu.com/web_api/sns/v3/search/note?keyword=中文&page=1&page_size=20", nil)
	wd := wr.GetContent("http://www.xiaohongshu.com/web_api/sns/v3/search/note?keyword=中文&page=1&page_size=20")
	for wd == nil {
		wd = wr.GetContent("http://www.xiaohongshu.com/web_api/sns/v3/search/note?keyword=中文&page=1&page_size=20")
		time.Sleep(3 * time.Second)
	}

	fmt.Println(wd.Result)
}

func TestWdRequest_SyncGet(t *testing.T) {
	wd := NewWdRequest(1)
	rs := wd.URLsIsComplete([]string{"https://www.baidu.com/s?wd=site%3Ajscrdq.com", "https://www.baidu.com/s?wd=site%3Ash-hting.com", "https://www.baidu.com/s?wd=site%3Ajustwin.cn", "https://www.baidu.com/s?wd=site%3Aalfaromeo.com.cn", "https://www.baidu.com/s?wd=site%3Aboe.com", "https://www.baidu.com/s?wd=site%3Aomegawatches.cn", "https://www.baidu.com/s?wd=site%3Aitaojin.cn", "https://www.baidu.com/s?wd=site%3Aozner.net", "https://www.baidu.com/s?wd=site%3Anobiliachina.com", "https://www.baidu.com/s?wd=site%3Asamsonite.com.cn"})
	for _, r := range rs {
		fmt.Println(r)
	}
	//wd.SyncGet("https://www.baidu.com/s?wd=%E7%8B%AC%E5%A2%85%E6%B9%96%E5%8C%BB%E9%99%A2&rn=50&pn=0", 10)
}
