//++++++++++++++++++++++++++++++++++++++++
//《Go语言高级开发与实战》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 知乎：https://www.zhihu.com/people/shirdonl
// 公众号:源码大数据
// 仓库地址：https://gitee.com/shirdonl/goAdvanced
// 仓库地址：https://github.com/shirdonl/goAdvanced
//++++++++++++++++++++++++++++++++++++++++

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func MySQLLock() (*sql.Tx, error) {
	db, _ = sql.Open("mysql",
		"root:a123456@tcp(127.0.0.1:3306)/chapter6")
	tx, err := db.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		fmt.Printf("SELECT * FROM methodLock WHERE `method_name`='method_name' FOR UPDATE", err)
		return nil, nil
	}
	return tx, err
}

func main() {
	res, _ := MySQLLock()
	//执行逻辑
	fmt.Println("执行逻辑")
	res.Commit() // 提交事务
	fmt.Println("exec transaction success!")
}
