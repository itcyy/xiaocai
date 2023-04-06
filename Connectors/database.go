package Connectors

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

/*
*   ConnMySQL
*   @author: [陈永裕]
*   @version[v1.0.0.1,2023-4-1]
*   @Description:
*   @return *sql.DB
 */
func ConnMySQL() *sql.DB {
	// 数据源名
	driverName := "mysql"
	// 用户名root
	// 密码1234
	// tcp协议连接
	// 数据库地址
	// 数据库 wms
	dataSourceName := "cyy" + ":" + "Chenyonyu1@" + "@" + "tcp" + "(" + "175.178.96.147:3306" + ")" + "/" + "data"
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}

	// 数据库设置
	db.SetConnMaxLifetime(time.Minute * 10)
	db.SetConnMaxIdleTime(time.Minute * 10)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// 连接测试
	err = db.Ping()
	if err != nil {
		fmt.Println("数据库连接失败")
		panic(err)
	}
	return db
}
