package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()
	if err != nil {
		// 错误处理逻辑
	}
	defer db.Close()
}
