package initDB

import (
	"database/sql"
	redis2 "github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB
var rdb *redis2.Client

func init() {
	initDb()
	//initRdb()
}

func initDb() {
	database, err := sql.Open("mysql", "jb5982857:wobuailq99@tcp(rm-2vc000gt370tk69pj4o.mysql.cn-chengdu.rds.aliyuncs.com:3306)/cat")
	if err != nil {
		log.Panicln("init mysql err:", err.Error())
	}
	Db = database
	database.SetMaxOpenConns(10)
	database.SetMaxIdleConns(10)
}

func initRdb() {
	rdb = redis2.NewClient(&redis2.Options{
		Addr:     "r-2vcvyp1wvczzixwrm8pd.redis.cn-chengdu.rds.aliyuncs.com:6379",
		Password: "wobuailq99JB", // no password set
		DB:       0,              // use default DB
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		log.Panicln("init redis err:", err.Error())
		return
	}
}
