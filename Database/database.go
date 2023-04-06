package Database

import (
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

//用户结构体
type Users struct {
	DataId    string `db:"id"`
	DataName  string `db:"Bookname"`
	Dataautor string `db:"autor"`
	DataTime  string `db:"time"`
	DataNum   string `db:"number"`
	DataUrl   string `db:"url"`
	Datadata  string `db:"data"`
}

//数据库指针
var db *sqlx.DB

//初始化数据库连接，init()方法系统会在动在main方法之前执行。
func init() {
	fmt.Printf("10086")
	database, err := sqlx.Open("mysql", "cyy:Chenyonyu1@@tcp(175.178.96.147:3306)/data")
	if err != nil {
		fmt.Println("open mysql failed,", err)
	}
	db = database
}

/*
*   SelectData
*   @author: [陈永裕]
*   @version[v1.0.0.1,2023-4-1]
*   @param [str,string]
*   @Description:
*   @param str
*   @return []Users
 */
func SelectData(str string) []Users {
	var user Users
	var fal []Users
	fmt.Printf("1005586")
	sql := str

	//执行SQL语句
	res, err := db.Query(sql)
	if err != nil {
		fmt.Println("exec failed,", err)
		return fal
	}
	var data []Users
	for res.Next() {
		err := res.Scan(&user.DataId, &user.DataName, &user.Dataautor, &user.DataTime, &user.DataNum, &user.DataUrl, &user.Datadata)
		if err != nil {
			log.Printf("scan failed, err:%v\n", err)
			return fal
		}
		data = append(data, user)

	}
	return data
}

/*
*   InsertData
*   @author: [陈永裕]
*   @version[v1.0.0.1,2023-4-1]
*   @param [sqlStr,string]
*   @param [username,string]
*   @param [password,string]
*   @param [email,string]
*   @Description:
*   @param sqlStr
*   @param username
*   @param password
*   @param email
*   @return bool
 */
func InsertData(sqlStr string, username string, password string, email string) bool {
	/*sqlStr := "insert into login(username,password,email) value(?,?,?)"*/
	_, err := db.Exec(sqlStr, username, password, email)
	if err != nil {
		fmt.Println("插入失败")
		log.Panic(err.Error())
		return false
	} else {
		return true
		fmt.Println("插入成功")
	}
	return false
}
