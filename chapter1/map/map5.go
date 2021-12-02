package main

import (
	"fmt"
	"sort"
)

type Programmer struct {
	FirstName string
	GoodAt    string
}

func main() {
	members := []Programmer{
		{"Jack", "PHP"},
		{"Jane", "JAVA"},
		{"Barry", "Go"},
	}

	fmt.Println(members)

	sort.Slice(members, func(i, j int) bool {
		return members[i].GoodAt < members[j].GoodAt || members[i].FirstName < members[j].FirstName
	})

	fmt.Println(members)
}
