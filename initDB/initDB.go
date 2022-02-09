package initDB

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

func init() {
	Db, err := sql.Open("mysql", "root:1234@tcp(rm-2vc000gt370tk69pj4o.mysql.cn-chengdu.rds.aliyuncs.com:3306)/cat")
	if err != nil {
		log.Panicln("err:", err.Error())
	}

	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)
}
