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
	var name string
	err = db.QueryRow("select name from user where id = ?", 6).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			//结果没有行，但没有发生错误
		} else {
			log.Fatal(err)
		}
	}
	fmt.Println(name)
}
