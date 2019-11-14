package mysql

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql" //import mysql impolementation for sql
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
)

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

//IDB 数据库接口
type IDB interface {
	GetDB() *sqlx.DB
	Prefix(str string) string
}

type sqlServer struct {
	Config Config
	DB     *sqlx.DB
}

//init 初始化sql
func (sql *sqlServer) init(config Config) error {
	var dbonce sync.Once
	var db *sqlx.DB
	var err error
	dbonce.Do(func() {
		db, err = sqlx.Open(
			"mysql",
			fmt.Sprintf(
				"%s:%s@tcp(%s:%d)/%s",
				config.User,
				config.Password,
				config.Host,
				config.Port,
				config.Db,
			),
		)
		if err != nil {
			log.Printf("get mysql database error: %s", err)
		} else {
			db.SetConnMaxLifetime(time.Duration(config.ConnLifeTime) * time.Second)
			db.SetMaxIdleConns(config.MaxIdleConn)
			db.SetMaxOpenConns(config.MaxOpenConn)
			db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)
		}
	})

	sql.DB = db
	return err
}

//GetDB GetDB
func (sql *sqlServer) GetDB() *sqlx.DB {
	return sql.DB
}

//Prefix change the relative sql to real sql with prefix
func (sql sqlServer) Prefix(str string) string {
	return strings.Replace(str, "#__", sql.Config.Dbprefix, -1)
}

var s IDB

//InitSQL 初始化数据库
func InitMYSQL(config Config) {
	ss := &sqlServer{
		Config: config,
	}
	err := ss.init(config)
	if err != nil {
		fmt.Println(err)
	}
	s = ss
}

//GetDB 获取数据库
func GetDB() *sqlx.DB {
	return s.GetDB()
}

//改sql表前缀
func Prefix(sql string) string {
	return s.Prefix(sql)
}
