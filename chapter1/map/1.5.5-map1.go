package main

import "fmt"

func main() {
	m := map[string]map[string]string{}
	m["programmer"] = map[string]string{}
	m["programmer"]["name"] = "Shirdon"
	m["manager"] = map[string]string{}
	m["manager"]["goodAt"] = "Go"
	fmt.Println(m["programmer"]["name"])
	fmt.Println(m)
	for key, value := range m {
		for k, v := range value {
			fmt.Println("K:", k, "V:", v)
		}
		fmt.Println("Key:", key, "Value:", value)
	}
}

//K: name V: Shirdon
//Key: programmer Value: map[name:Shirdon]
//K: goodAt V: Go
//Key: manager Value: map[goodAt:Go]
