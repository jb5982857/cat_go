package initDB

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

func init() {
	database, err := sql.Open("mysql", "jb5982857:wobuailq99@tcp(rm-2vc000gt370tk69pj4o.mysql.cn-chengdu.rds.aliyuncs.com:3306)/cat")
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	Db = database
	database.SetMaxOpenConns(10)
	database.SetMaxIdleConns(10)
}
