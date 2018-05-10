package mysqlutil

import (
	"io/ioutil"
	"encoding/json"
)


type ProjectDbManager struct {
	MySQLServers []MySQLServer `json:"mysql_servers"`
	DefaultServerName string `json:"default_server_name"`
}


type MySQLServer struct {
	ServerName string `json:"server_name,omitempty"`
	UserName string `json:"user_name"`
	Host string `json:"host"`
	Port int `json:"port"`
	PassWord string `json:"pass_word"`
	Databases []string `json:"databases"`
	dbMus map[string]MysqlUtil
}

// 从文件中读取配置
func (pdm *ProjectDbManager) InitServers(confPath string) error{
	jsonData,err := ioutil.ReadFile(confPath)
	if err != nil {
		return err
	}
	json.Unmarshal(jsonData,pdm)
	return nil
}

// 根据server的名称和数据库名获取对应的db
func (pdm *ProjectDbManager) GetDb(serverName string, dbName string) (MysqlUtil, error) {
	msInstance := MysqlUtil{}
	for _, mySQLServer := range pdm.MySQLServers{
		if mySQLServer.ServerName == serverName{
			if mysqlDb,ok := mySQLServer.dbMus[dbName];ok{
				return mysqlDb,nil
			}
			if mySQLServer.dbMus == nil {
				mySQLServer.dbMus = make(map[string]MysqlUtil)
			}
			mySQLServer.dbMus[dbName] = msInstance
			err := msInstance.InitMySqlUtilDetail(mySQLServer.Host, mySQLServer.Port, mySQLServer.UserName, mySQLServer.PassWord,dbName,1,2)
			if err != nil {
				return msInstance,err
			}
			return msInstance, nil
		}
	}
	return msInstance,nil
}

// 获取默认的db
func (pdm *ProjectDbManager) GetDefaultServerDb(DbName string) (MysqlUtil, error) {
	return pdm.GetDb(pdm.DefaultServerName, DbName)
}