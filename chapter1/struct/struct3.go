package main

import (
	"encoding/json"
	"fmt"
)

//基础结构体
type Item struct {
	Title string
	URL   string
}

//定义一个响应结构体
type Response struct {
	Data struct {
		Children []struct {
			Data Item
		}
	}
}

func main() {
	jsonStr := `{
	"data": {
		"children": [
			{
				"data": {
					"title": "Shirdon's Blog'",
					"url": "https://www.shirdon.com"
				}
			}
		]
	}
}`
	res := Response{}
	json.Unmarshal([]byte(jsonStr), &res)
	fmt.Println(res)
}

//{{[{{Shirdon's Blog' https://www.shirdon.com}}]}}
