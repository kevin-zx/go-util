package mysqlutil


import (
"database/sql"
"fmt"
_ "github.com/go-sql-driver/mysql"
)

type MysqlUtil struct {
	db *sql.DB
	
}
var GlobalMysqlUtil MysqlUtil

func (mu *MysqlUtil) initMySqlUtil(host string, port int, user string, passwd string, databases string, maxIdleConns int,MaxOpenConns int) error {
	// 避免重复调用
	if mu.db != nil{
		mu.Close()
	}
	dataSourceNameFormat := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",user,passwd,host,port,databases)
	println(dataSourceNameFormat)
	mu.db, _ = sql.Open("mysql", dataSourceNameFormat)
	mu.db.SetMaxIdleConns(maxIdleConns)
	mu.db.SetMaxOpenConns(MaxOpenConns)
	//mu.db.Close()
	err := mu.db.Ping()
	return err
}
func (mu *MysqlUtil) InitMySqlUtilByDb(db *sql.DB) error {
	mu.db = db
	err := db.Ping()
	return err
}

// 判断是否初始化
func (mu *MysqlUtil) IsInit() bool {
	return mu.db != nil
}

func (mu *MysqlUtil) IsAlive() bool  {
	err := mu.db.Ping()
	if err != nil {
		mu.Close()
		return false
	}
	return true
}

func (mu *MysqlUtil) Close() error {
	err := mu.db.Close()
	return err
}

func (mu *MysqlUtil) InitMySqlUtilDetail(host string, port int, user string, passwd string, databases string, axIdleConns int,MaxOpenConns int) error {
	return mu.initMySqlUtil(host,port,user,passwd,databases,axIdleConns,MaxOpenConns)
}

func (mu *MysqlUtil) InitMySqlUtil(host string, port int, user string, passwd string, databases string) error {
	return mu.initMySqlUtil(host,port,user,passwd,databases,0,1)
}


func (mu *MysqlUtil) Insert(prepareSql string, args ...interface{}) error {
	stmt, err := mu.db.Prepare(prepareSql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(args...)
	if err != nil {
		return err
	}
	return nil
}

func (mu *MysqlUtil) InsertId(prepareSql string, args ...interface{}) (int64,error) {
	stmt, err := mu.db.Prepare(prepareSql)
	if err != nil {
		return 0,err
	}
	defer stmt.Close()
	re, err := stmt.Exec(args...)
	if err != nil {
		return 0,err
	}
	lastId,err := re.LastInsertId()
	if err != nil {
		return 0,err
	}
	return lastId,nil
}

func (mu *MysqlUtil) Exec(prepareSql string, args ...interface{}) error {
	stmt, err := mu.db.Prepare(prepareSql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(args...)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}




func (mu *MysqlUtil) ExecBatch(prepareSql string, args [][]interface{}) error {
	tx, err := mu.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare(prepareSql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	for _, arg := range args {
		_, err = stmt.Exec(arg...)
		if err != nil {
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil

}

func (mu *MysqlUtil) Select(prepareSql string, args ...interface{}) ([][]sql.RawBytes, error) {
	rows, err := mu.db.Query(prepareSql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	var valueArr [][]sql.RawBytes
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}
		var value = make([]sql.RawBytes, len(columns))
		copy(value, values[:])
		valueArr = append(valueArr, value)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return valueArr, nil
}


func (mu *MysqlUtil) SelectAll(sqlstr string, args ...interface{}) (*[]map[string]string, error) {
	stmtOut, err := mu.db.Prepare(sqlstr)
	if err != nil {
		//panic(err.Error())
		return nil,err
	}
	defer stmtOut.Close()

	rows, err := stmtOut.Query(args...)
	if err != nil {
		//panic(err.Error())
		return nil,err
	}

	columns, err := rows.Columns()
	if err != nil {
		//panic(err.Error())
		return nil,err
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))

	ret := make([]map[string]string, 0)
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			//panic(err.Error())
			return nil,err
		}
		var value string
		vmap := make(map[string]string, len(scanArgs))
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			vmap[columns[i]] = value
		}
		ret = append(ret, vmap)
	}
	return &ret, nil
}
