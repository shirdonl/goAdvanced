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
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/down", Welcome)
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	//定义导出的文件名
	filename := "exportUsers.csv"

	//定义一个二维数组
	column := [][]string{
		{"手机号", "用户UID", "Email", "用户名"},
		{"18888888888", "2", "barry@163.com", "barry"},
		{"18888888889", "3", "wangwu@163.com", "wangwu"},
	}

	//导出
	GenerateCsv(filename, column)
	//读取文件
	file, err := os.Open(filename)
	content, err := ioutil.ReadAll(file)
	//下载文件
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("content-disPosition", "attachment; filename=\""+filename+"\"")
	if err != nil {
		fmt.Println("Read File Err:", err.Error())
	} else {
		w.Write(content)
	}
}

//生成CSV文件
func GenerateCsv(filePath string, data [][]string) {
	fp, err := os.Create(filePath) //创建文件句柄
	if err != nil {
		log.Fatalf("创建文件["+filePath+"]句柄失败,%v", err)
		return
	}
	defer fp.Close()
	fp.WriteString("\xEF\xBB\xBF") //写入UTF-8 BOM
	w := csv.NewWriter(fp)         //创建一个新的写入文件流
	w.WriteAll(data)
	w.Flush()
}
