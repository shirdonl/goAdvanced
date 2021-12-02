package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:a123456@tcp(127.0.0.1:3306)/ch4")
	if err != nil {
		fmt.Println(err)
	}
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	var (
		id    int
		phone string
	)
	//通过defer语句回滚
	defer tx.Rollback()
	//查询数据库
	rows, err := db.Query("select id, phone from user where id = ?", 3)
	if err != nil {
		log.Fatal(err)
	}
	//defer语句关闭数据库连接
	defer rows.Close()
	for rows.Next() {
		//通过Scan方法赋值
		err := rows.Scan(&id, &phone)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, phone)
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
