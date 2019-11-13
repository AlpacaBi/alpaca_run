package mysql

import (
	_ "github.com/go-sql-driver/mysql" //import mysql impolementation for sql
	"github.com/jmoiron/sqlx"
	"hz.com/golib/mysql"
)

var s mysql.IDB

//Config Mysql的配置
type Config struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	User         string `json:"user"`
	Password     string `json:"password"`
	Db           string `json:"db"`
	Dbprefix     string `json:"dbprefix"`
	ConnLifeTime int    `json:"connlifetime"` //以秒为单位
	MaxIdleConn  int    `json:"maxidleconn"`
	MaxOpenConn  int    `json:"maxopenconn"`
}

//InitSQL 初始化数据库
func InitMYSQL(config mysql.Config) {
	ss, er := mysql.New(config)
	if er != nil {
		panic(er)
	}
	s = ss
}

//GetDB 获取数据库
func GetDB() *sqlx.DB {
	return s.GetDB()
}

//Prefix change the relative sql to real sql with prefix
func Prefix(sql string) string {
	return s.Prefix(sql)
}

//UnPrefix change the real sql with prefix to relative one
func UnPrefix(sql string) string {
	return s.UnPrefix(sql)
}
