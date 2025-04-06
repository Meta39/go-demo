package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

/*
DB 定义一个全局对象
sql.DB是表示连接的数据库对象（结构体实例），它保存了连接数据库相关的所有信息。
它内部维护着一个具有零到多个底层连接的连接池，它可以安全地被多个goroutine同时使用。
SetMaxOpenConns
func (db *DB) SetMaxOpenConns(n int)
SetMaxOpenConns设置与数据库建立连接的最大数目。 如果n大于0且小于最大闲置连接数，会将最大闲置连接数减小到匹配最大开启连接数的限制。如果n<=0，不会限制最大开启连接数，默认为0（无限制）。

SetMaxIdleConns
func (db *DB) SetMaxIdleConns(n int)
SetMaxIdleConns设置连接池中的最大闲置连接数。 如果n大于最大开启连接数，则新的最大闲置连接数会减小到匹配最大开启连接数的限制。 如果n<=0，不会保留闲置连接。
*/
var DB *sql.DB

// InitDB 定义一个初始化数据库的函数
func InitDB() (err error) {
	//数据库建表语句：go_mysql_demo.sql
	dsn := "root:123456@tcp(localhost:3306)/go_mysql_demo?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// 连接池配置
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)

	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	DB = db
	return nil
}
