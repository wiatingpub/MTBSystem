package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 数据库实例
var (
	db *sqlx.DB
)

// Init 初始化
func Init(mysqlDSN string) {
	db = sqlx.MustConnect("mysql", mysqlDSN)
	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(3)
}
