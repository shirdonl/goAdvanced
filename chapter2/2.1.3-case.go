package main

// 从Mysql中导出数据到CSV文件
import (
	"database/sql"
	"encoding/csv"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var (
	tables = []string{"user", "order"}
	count  = len(tables)
	ch     = make(chan bool, count)
)

func main() {
	db, err := sql.Open("mysql", "root:a123456@tcp(127.0.0.1:3306)/chapter4?charset=utf8")
	defer db.Close()

	if err != nil {
		panic(err.Error())
	}

	for _, table := range tables {
		go SqlQuery(db, table, ch)
	}

	for i := 0; i < count; i++ {
		<-ch
	}
	fmt.Println("完成!")
}

//运行SQL
func SqlQuery(db *sql.DB, table string, ch chan bool) {
	fmt.Println("开始处理：", table)
	rows, _ := db.Query(fmt.Sprintf("SELECT * from %s", table))

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	totalValues := [][]string{}
	for rows.Next() {
		var s []string
		//把每行的内容添加到scanArgs，也添加到了values
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		for _, v := range values {
			s = append(s, string(v))
		}
		totalValues = append(totalValues, s)
	}

	if err = rows.Err(); err != nil {
		panic(err.Error())
	}
	exportToCSV(table+".csv", columns, totalValues)
	ch <- true
}

//导出到CSV
func exportToCSV(file string, columns []string, totalValues [][]string) {
	f, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	f.WriteString("\xEF\xBB\xBF")
	defer f.Close()
	w := csv.NewWriter(f)
	for a, i := range totalValues {
		if a == 0 {
			w.Write(columns)
			w.Write(i)
		} else {
			// fmt.Println(i)
			w.Write(i)
		}
	}
	w.Flush()
	fmt.Println("处理完毕：", file)
}
