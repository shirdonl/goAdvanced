package main

import (
	"encoding/csv"
	"os"
)

func main() {
	//创建一个名为"test.csv"的文件
	f, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM

	//NewWriter()函数返回一个Writer对象
	w := csv.NewWriter(f)
	//调用Writer对象的Write()方法写入数据到CSV文件
	w.Write([]string{"学号", "姓名", "分数"})
	w.Write([]string{"1", "Barry", "99.5"})
	w.Write([]string{"2", "Shirdon", "100"})
	w.Write([]string{"3", "Jack", "88"})
	w.Write([]string{"4", "Dong", "68"})
	w.Flush()
}
