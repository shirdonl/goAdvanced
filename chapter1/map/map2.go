package main

import (
	"encoding/json"
	"fmt"
	"github.com/goinggo/mapstructure"
	"testing"
)

func main() {
	//jsonStr := `
	//   {
	//           "name": "Shirdon",
	//           "goodAt": "Go Programming"
	//   }
	//   `
	//var mapResult map[string]interface{}
	//err := json.Unmarshal([]byte(jsonStr), &mapResult)
	//if err != nil {
	//	fmt.Println("JsonToMapDemo err: ", err)
	//}
	//fmt.Println(mapResult)

	instance := map[string]interface{}{
		"name":   "Shirdon",
		"goodAt": "Go Programming"}

	jsonStr, err := json.Marshal(instance)

	if err != nil {
		fmt.Println("MapToJsonDemo err: ", err)
	}
	fmt.Println(string(jsonStr))
}

func TestMapToStruct(t *testing.T) {
	mapInstance := make(map[string]interface{})
	mapInstance["Name"] = "liang637210"
	mapInstance["Age"] = 28

	var people People
	//将 map 转换为指定的结构体
	if err := mapstructure.Decode(mapInstance, &people); err != nil {
		t.Fatal(err)
	}
	t.Logf("map2struct后得到的 struct 内容为:%v", people)
}
