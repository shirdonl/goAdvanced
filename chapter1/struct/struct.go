package main

import (
	"fmt"
	"sort"
)

type Programmer struct {
	Name string
	Age  int
}

type ProgrammerSlice []Programmer

func (s ProgrammerSlice) Len() int           { return len(s) }
func (s ProgrammerSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ProgrammerSlice) Less(i, j int) bool { return s[i].Age < s[j].Age }

func main() {
	a := ProgrammerSlice{
		{
			Name: "Barry",
			Age:  30,
		},
		{
			Name: "Jack",
			Age:  22,
		},
		{
			Name: "Jim",
			Age:  18,
		},
	}
	sort.Stable(a)
	fmt.Println(a)
}

//[{Jim 18} {Jack 22} {Barry 30}]
