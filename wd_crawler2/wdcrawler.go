package wd_crawler

import (
	"encoding/base64"
	"encoding/json"
	"github.com/kevin-zx/go-util/dateUtil"
	"github.com/kevin-zx/go-util/mysqlUtil"
	"strconv"
	"sync"
	"time"
	//"sync"
	"log"
)

type WdRequest struct {
	Header        map[string]string
	mu            mysqlutil.MysqlUtil
	mux           sync.Mutex
	Redirect      int
	Single        int
	ConcurrentNum int
	StoreType     int
	ContentType   int
	isLan         bool
	mysqlConfig   mysqlConfig
}

type mysqlConfig struct {
	mysqlHost    string
	mysqlPort    int
	mysqlUser    string
	mysqlPasswd  string
	mysqlDb      string
	axIdleConns  int
	maxOpenConns int
}

type WdResponse struct {
	Header      string
	Result      string
	Code        int
	RedirectUrl string
	Status      int
}

func (wc *WdRequest) SendRequest(targetUrl string, header *map[string]string) error {
	if header == nil {
		header = &wc.Header
	}
	if (*header)["User-Agent"] == "" && (*header)["user-agent"] == "" {
		(*header)["User-Agent"] = wc.Header["User-Agent"]
	}
	headerJson, _ := json.Marshal(header)

	configInsertSql := "INSERT INTO configs_16 (`header`,`redirect`,`single`,`concurrent_num`,`store_type`,`param`,`expire_time`)" +
		" VALUE (?,?,?,?,?,?,?)"

	id, err := wc.mu.InsertId(configInsertSql, headerJson, wc.Redirect, wc.Single, wc.ConcurrentNum, wc.StoreType, "", dateUtil.GetDeltaDateTime(2*time.Hour))
	if err != nil {
		//panic(err)
		return err
	}

	urlInsertSql := "INSERT IGNORE INTO urls_16 (`url`,`md5`,`type`,`config_id`,`status`) VALUE (?,md5(?),?,?,?)"
	err = wc.mu.Exec(urlInsertSql, targetUrl, targetUrl, wc.ContentType, id, 0)
	if err != nil {
		return err
	}
	return nil

}

func (wc *WdRequest) ExistURL(targetUrl string) (bool, error) {
	data, err := wc.mu.SelectAll("SELECT * FROM urls_16 where `md5` = md5(?) AND type=? AND status !=3 LIMIT 1", targetUrl, wc.ContentType)
	if err != nil {
		//return false,err
		log.Println(err)
	}
	return data != nil && len(*data) > 0, err
}

//port 1 pc 2 mobile
func NewWdRequest(port int, mysqlHost string, mysqlPort int, mysqlUser string, mysqlPasswd string, mysqlDb string, axIdleConns int, maxOpenConns int) *WdRequest {
	mysqlConfig := mysqlConfig{mysqlHost: mysqlHost, mysqlPort: mysqlPort, mysqlUser: mysqlUser, mysqlPasswd: mysqlPasswd, mysqlDb: mysqlDb, axIdleConns: axIdleConns, maxOpenConns: maxOpenConns}
	header := make(map[string]string)
	mu := mysqlutil.MysqlUtil{}
	err := mu.InitMySqlUtilDetail(mysqlConfig.mysqlHost, mysqlConfig.mysqlPort, mysqlConfig.mysqlUser, mysqlConfig.mysqlPasswd, mysqlConfig.mysqlDb, mysqlConfig.axIdleConns, mysqlConfig.maxOpenConns)
	if err != nil {
		panic(err)
	}
	if port == 1 {
		header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.139 Safari/537.36"
	} else {
		header["User-Agent"] = "Mozilla/5.0 (Linux; Android 5.0; SM-G900P Build/LRX21T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.139 Mobile Safari/537.36"
	}
	return &WdRequest{Header: header, mu: mu, Redirect: 1, Single: 0, ConcurrentNum: 0, StoreType: 1, ContentType: 1, isLan: false, mysqlConfig: mysqlConfig}
}

func (wc *WdRequest) RestMysqlConnection() error {
	var err error
	if wc.mu.IsAlive() {
		return err
	}
	wc.mu.Close()

	mu := mysqlutil.MysqlUtil{}
	mysqlConfig := wc.mysqlConfig
	err = mu.InitMySqlUtilDetail(mysqlConfig.mysqlHost, mysqlConfig.mysqlPort, mysqlConfig.mysqlUser, mysqlConfig.mysqlPasswd, mysqlConfig.mysqlDb, mysqlConfig.axIdleConns, mysqlConfig.maxOpenConns)
	if err != nil {
		return err
	}
	wc.mu = mu

	return nil
}

func (wc *WdRequest) Close() {
	wc.mu.Close()
}

func (wc *WdRequest) DeleteUrlTask(targetUrl string, status int) {
	if status == -1 {
		wc.mu.Exec("DELETE FROM urls_16 where `md5` = md5(?) AND type=? AND status = ?", targetUrl, wc.ContentType, status)
	} else {
		wc.mu.Exec("DELETE FROM urls_16 where `md5` = md5(?) AND type=?", targetUrl, wc.ContentType)
	}
}

func (wc *WdRequest) GetContent(targetUrl string) *WdResponse {
	data, err := wc.mu.SelectAll("SELECT result,header,code,redirect_url,status FROM urls_16 WHERE `md5` = md5(?) AND `status` > 1 AND type=?", targetUrl, wc.ContentType)
	if err != nil {
		wc.RestMysqlConnection()
		return nil
	}
	if len(*data) == 0 {
		return nil
	}
	d := (*data)[0]
	baseResult := d["result"]
	reByte, _ := base64.StdEncoding.DecodeString(baseResult)
	code, _ := strconv.Atoi(d["code"])
	redirect := d["redirect_url"]
	status, _ := strconv.Atoi(d["status"])
	wr := WdResponse{Result: string(reByte), Code: code, RedirectUrl: redirect, Status: status}
	if status == 3 {
		//如果任务错误了 顺手删除了
		wc.mu.Exec("DELETE FROM urls_16 where `md5` = md5(?) AND type=? AND status = 3", targetUrl, wc.ContentType)
	}
	return &wr
}

// sysc get content, every retryTime will sleeping 1 min
// 同步的获取内容，retryTime 代表重试次数，每次会等待1分钟
func (wc *WdRequest) SyncGet(targetUrl string, retryTime int) (*WdResponse, error) {
	return wc.SyncGetWithHeader(targetUrl, nil, retryTime)
}

// sysc get content, every retryTime will sleeping 1 min
// 同步的获取内容，retryTime 代表重试次数，每次会等待1分钟, header 是请求头
func (wc *WdRequest) SyncGetWithHeader(targetUrl string, header map[string]string, retryTime int) (*WdResponse, error) {
	//HACK: 这里的实现方式非常粗糙 todo:换一种实现方式
	exist, err := wc.ExistURL(targetUrl)

	if err != nil {
		time.Sleep(3 * time.Second)
		wc.RestMysqlConnection()
		exist, err = wc.ExistURL(targetUrl)
	}
	if err != nil {
		return nil, err

	}
	if !exist {
		wc.SendRequest(targetUrl, &header)
		time.Sleep(30 * time.Second)
	}

	wd := wc.GetContent(targetUrl)
	for i := 0; wd == nil && i < retryTime*6; i++ {
		time.Sleep(10 * time.Second)
		wd = wc.GetContent(targetUrl)
	}
	return wd, nil
}
