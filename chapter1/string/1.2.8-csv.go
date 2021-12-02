package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./test.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	reader.FieldsPerRecord = -1

	for {
		record, e := reader.Read()
		if e != nil {
			fmt.Println(e)
			break
		}
		fmt.Println(record)
	}
}

//[手机号 用户UID Email 用户名]
//[18888888888 2 barry@163.com barry]
//[18888888889 3 wangwu@163.com wangwu]
//EOF
