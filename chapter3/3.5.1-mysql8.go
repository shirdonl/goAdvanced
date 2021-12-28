//++++++++++++++++++++++++++++++++++++++++
// 《Go语言高级开发与实战》源码
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
	"log"
	"time"
)

func main() {
	db, err := sql.Open("mysql", "root:a123456@tcp(127.0.0.1:3306)/ch4")
	if err != nil {
		fmt.Println(err)
	}
	go func(db *sql.DB) {
		mt := time.NewTicker(10 * time.Second)
		for {
			select {
			case <-mt.C:
				stat := db.Stats()
				fmt.Println(stat.MaxOpenConnections)
				fmt.Errorf("monitor db conn(%p): maxopen(%d), open(%d), use(%d), idle(%d), "+
					"wait(%d), idleClose(%d), lifeClose(%d), totalWait(%v)",
					db,
					stat.MaxOpenConnections, stat.OpenConnections,
					stat.InUse, stat.Idle,
					stat.WaitCount, stat.MaxIdleClosed,
					stat.MaxLifetimeClosed, stat.WaitDuration)
			}
		}
	}(db)
	stmt, err := db.Prepare("select id, phone from user where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(3)
	for rows.Next() {
		var s sql.NullString
		err = rows.Scan(&s)
		fmt.Println(s)
		// 检查错误
		if s.Valid {
			// 用 s.String来获取值
		} else {
			// NULL值
		}
	}

}
