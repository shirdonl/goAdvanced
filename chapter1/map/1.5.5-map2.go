package main

import (
	"fmt"
)

func main() {
	/* 使用interface{}初始化一个一维映射
	 * 关键点：interface{} 可以代表任意类型
	 * 原理知识点：interface{} 就是一个空接口，所有类型都实现了这个接口，所以它可以代表所有类型
	 */
	//mainMap := make(map[string]interface{})
	mainMap := map[string]interface{}{}
	subMap := make(map[string]string)
	subMap["Key_1"] = "SubValue_1"
	subMap["Key_2"] = "SubValue_2"
	mainMap["Map"] = subMap
	for key, val := range mainMap {
		//此处必须实例化接口类型，即*.(map[string]string)
		//subMap := val.(map[string]string)
		for subKey, subVal := range val.(map[string]string) {
			fmt.Printf("mapName=%s    Key=%s    Value=%s\n", key, subKey, subVal)
		}
	}
}

//mapName=Map    Key=Key_1    Value=SubValue_1
//mapName=Map    Key=Key_2    Value=SubValue_2
