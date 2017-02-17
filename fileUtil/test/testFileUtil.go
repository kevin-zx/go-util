package main

import (
	"github.com/kevin-zx/go-util/fileUtil"
	"strings"
	//"github.com/kevin-zx/go-util/mysqlUtil"
	"github.com/kevin-zx/go-util/mysqlUtil"
	"github.com/kevin-zx/go-util/regexpUtil"
)

func main()  {
	mysqlutil.GlobalMysqlUtil.InitMySqlUtil("115.159.3.51",3306, "remote","Iknowthat","eb_bigdata")
	c_data := fileUtil.ReadFile("C:/Users/Administrator/Desktop/文件/电商/分类 - 副本.txt","GBK")
	for _,c := range strings.Split(c_data, "\n"){
		cat_info := regexpUtil.SplitString(c,",| >> ")
		if len(cat_info) == 1{
			continue
		}
		var a []interface{}
		for _,ci := range(cat_info) {
			a = append(a,ci)
		}
		for len(a)<8  {
			a = append(a,"")
		}
		mysqlutil.GlobalMysqlUtil.Insert("INSERT INTO cate_tmp " +
			"(`cate_id`,`cate_name`,`level`,`level_1_name`,`level_2_name`,`level_3_name`,`level_4_name`,`level_5_name`)" +
			"value (?,?,?,?,?,?,?,?)", a ...)
	}

}

