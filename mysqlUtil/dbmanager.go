package mysqlutil

import (
	"io/ioutil"
	"encoding/json"
	"github.com/kevin-zx/go-util/mysqlUtil"
)


type ProjectDbManager struct {
	MySQLServers []MySQLServer `json:"mysql_servers"`
	DefaultServerName string
}


type MySQLServer struct {
	ServerName string `json:"server_name,omitempty"`
	UserName string `json:"user_name"`
	Host string `json:"host"`
	Port int `json:"port"`
	PassWord string `json:"pass_word"`
	Databases []string `json:"databases"`
	dbMus map[string]mysqlutil.MysqlUtil
}

// 从文件中读取配置
func (pdm *ProjectDbManager) InitServers(confPath string) error{
	json_data,err := ioutil.ReadFile(confPath)
	if err == nil {
		return err
	}
	json.Unmarshal(json_data,pdm)
	return nil
}

// 根据server的名称和数据库名获取对应的db
func (pdm *ProjectDbManager) GetDb(serverName string, dbName string) (mysqlutil.MysqlUtil, error) {
	ms_instance := mysqlutil.MysqlUtil{}
	for _, mySQLServer := range pdm.MySQLServers{
		if mySQLServer.ServerName == serverName{
			if mysqlDb,ok := mySQLServer.dbMus[dbName];ok{
				return mysqlDb,nil
			}
			err := mysqlutil.MysqlUtil{}.InitMySqlUtilDetail(mySQLServer.Host, mySQLServer.Port, mySQLServer.UserName, mySQLServer.PassWord,dbName,1,2)
			if err != nil {
				return ms_instance,err
			}
			mySQLServer.dbMus[dbName] = ms_instance
			return ms_instance, nil
		}
	}
	return ms_instance,nil
}

// 获取默认的db
func (pdm *ProjectDbManager) GetDefaultServerDb(DbName string) (mysqlutil.MysqlUtil, error) {
	return pdm.GetDb(pdm.DefaultServerName, DbName)
}