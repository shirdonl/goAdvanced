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
