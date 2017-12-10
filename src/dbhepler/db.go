/*
这是一个数据库辅助代码，建议使用该demo的时候运行一下文件，当然了运行之前数据库的账户和密码要做相应的修改。
*/
package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"math/rand"
	"share/config"
	"time"
)

var schema = `
CREATE TABLE IF NOT EXISTS user (
    id INT UNSIGNED AUTO_INCREMENT,
    name VARCHAR(20),
 	address VARCHAR(20),
    phone VARCHAR(15),
	PRIMARY KEY (id)
)`

type User struct {
	id      int32  `db:"id"`
	name    string `DB:"NAME"`
	Address string `db:"address"`
	Phone   string `db:"phone"`
}

func main() {
	// 同时打开并连接数据库，返回错误信息
	db, err := sqlx.Connect("mysql", config.MysqlDSN)
	if err != nil {
		log.Fatalln(err)
	}

	// exec the schema or fail; multi-statement Exec behavior varies between
	// database drivers;  pq will exec them all, sqlite3 won't, ymmv
	db.MustExec(schema)

	tx := db.MustBegin()

	rand.Seed(time.Now().UnixNano())

	tx.MustExec("INSERT INTO user (id, name, address,phone) VALUES (?, ?, ? ,?)", nil, "ricoder"+GetRandomString(2), "guangdong lufen"+GetRandomString(5), "188141284"+GetRandomString(2))

	tx.Commit()

}

func GetRandomString(leng int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < leng; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
