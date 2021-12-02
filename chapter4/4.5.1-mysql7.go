package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		fmt.Println(err)
	}
	stmt, err := db.Prepare("select id, name from users where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(1)
	for rows.Next() {
		var s sql.NullString
		err = rows.Scan(&s)
		// 检查错误
		if s.Valid {
			// 用 s.String来获取值
		} else {
			// NULL值
		}
	}
}
