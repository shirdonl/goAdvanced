package main

import "fmt"

func main() {

	//声明并初始化一个map，key是int64类型，value是string类型
	varMap := make(map[int64]string)
	varMap[1] = "Go"
	varMap[2] = "Advanced"

	//声明一个int64数组，然后遍历数组，num是数组中的元素
	for _, num := range []int64{1, 2, 3, 4} {
		if _, ok := varMap[num]; ok {
			fmt.Printf("varMap中包含key:%d \n", num)
		} else {
			fmt.Printf("varMap中不包含key:%d\n", num)
		}
	}

	fmt.Println("_______分割线_______")

	for _, num := range []int64{1, 2, 3, 4} {
		//如果包含key，想知道value，就把返回值赋给一个变量，这儿用变量v
		if v, s := varMap[num]; s {
			fmt.Printf("varMap中包含key:%d,value值为:%s\n", num, v)
		} else {
			fmt.Printf("varMap中不包含key:%d\n", num)
		}
	}
}
