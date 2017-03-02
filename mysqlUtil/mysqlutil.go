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

func (mu *MysqlUtil) initMySqlUtil(host string, port int, user string, passwd string, databases string, maxIdleConns int,MaxOpenConns int)  {
	dataSourceNameFormat := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",user,passwd,host,port,databases)
	//println(dataSourceNameFormat)
	mu.db, _ = sql.Open("mysql", dataSourceNameFormat)
	mu.db.SetMaxIdleConns(maxIdleConns)
	mu.db.SetMaxOpenConns(MaxOpenConns)
	err := mu.db.Ping()
	if err != nil {
		panic(err)
	}

}

func (mu *MysqlUtil) InitMySqlUtil(host string, port int, user string, passwd string, databases string)  {
	mu.initMySqlUtil(host,port,user,passwd,databases,1,2)
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

func (mu *MysqlUtil) Exec(prepareSql string, args ...interface{}) error {
	stmt, err := mu.db.Prepare(prepareSql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(args...)
	fmt.Println(err)
	if err != nil {
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
		panic(err.Error())
	}
	defer stmtOut.Close()

	rows, err := stmtOut.Query(args...)
	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
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
			panic(err.Error())
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
