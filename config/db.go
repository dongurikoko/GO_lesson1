package config

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	// blank import for MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

// Driver名
const driverName = "mysql"

var conn *sql.DB

func init() {
	/* ===== データベースへ接続する. ===== */
	// ユーザ
	user := os.Getenv("MYSQL_USER")
	// パスワード
	password := os.Getenv("MYSQL_PASSWORD")
	// 接続先ホスト
	host := os.Getenv("MYSQL_HOST")
	// 接続先ポート
	port := os.Getenv("MYSQL_PORT")
	// 接続先データベース
	database := os.Getenv("MYSQL_DATABASE")

	// 接続情報は以下のように指定する.
	// user:password@tcp(host:port)/database
	var err error
	conn, err = sql.Open(driverName,
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, database))
	if err != nil {
		log.Fatal(err)
	}
	if err := conn.Ping(); err != nil {
		log.Fatalf("can't connect to mysql server. "+
			"MYSQL_USER=%s, "+
			"MYSQL_PASSWORD=%s, "+
			"MYSQL_HOST=%s, "+
			"MYSQL_PORT=%s, "+
			"MYSQL_DATABASE=%s, "+
			"error=%+v",
			user, password, host, port, database, err)
	}
}

func GetConn() (*sql.DB, error) { // Mysqlに接続されていたらconnを返す。接続されていなかったらエラーを返す
	if conn == nil {
		return nil, errors.New("failed to connect Mysql in GetConn")
	}
	return conn, nil
}
