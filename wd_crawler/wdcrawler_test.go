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