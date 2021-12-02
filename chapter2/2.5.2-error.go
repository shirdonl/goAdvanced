package main

import (
	"errors"
	"fmt"
	"runtime/debug"
)

var ERROR_NO_RECORD = errors.New("no record")
var ERROR_UNEXPECTED_EOF = errors.New("unexpected EOF")

type User struct {
	UserName string
}

//检测用户类型
//func (user *User) CheckUserType(userType string) ERROR {
//	switch userType {
//	case "normal":
//		return nil
//	case "vip":
//		return nil
//	}
//	return ERRORs.New("CheckUserType Error:" + userType)
//}

//检测用户类型
func (user *User) CheckUserType(userType string) bool {
	return userType == "normal" || userType == "vip"
}

//func (user *User) setUserName() ERROR {
//	user.UserName = "Shirdon"
//	return nil
//}

func (user *User) setUserName() {
	user.UserName = "Shirdon"
}

func main() {
	//user := User{}
	//user.setUserName()
	//
	//url := "https://www.baidu.com"
	//resp, err := http.Get(url)
	//if err != nil {
	//	return resp, err
	//}
	//
	//value, ok := cache.Lookup(key)
	//if !ok {
	//	// ...cache[key] does not exist…
	//}

	//err := user.setUserName()
	//if err != nil {
	//	// 返回错误
	//	return ERRORs.New("ERROR")
	//}
	//	funcOne()
	testPanic()

}

//func funcOne() error {
//	defer func() {
//		if p := recover(); p != nil {
//			fmt.Printf("异常恢复!p: %v", p)
//			debug.PrintStack()
//		}
//	}()
//	return funcTwo()
//}
//func funcTwo() error {
//	// 模拟
//	panic("模拟异常")
//	return errors.New("success")
//}
//
//func testPanic() {
//	err := funcOne()
//	if err == nil {
//		fmt.Printf("error is nil \n")
//	} else {
//		fmt.Printf("error is %v \n", err)
//	}
//}

//error is nil

func testPanic() {
	err := funcOne()
	if err == nil {
		fmt.Printf("error is nil \n")
	} else {
		fmt.Printf("error is %v \n", err)
	}
}

//
func funcOne() (err error) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("异常恢复!p:", p)
			str, ok := p.(string)
			if ok {
				err = errors.New(str)
			} else {
				err = errors.New("success")
			}
			debug.PrintStack()
		}
	}()
	return funcTwo()
}

func funcTwo() error {
	// 模拟
	panic("模拟异常")
	return errors.New("success")
}

//error is 模拟异常

func getSeasons() []string {
	return[]string{"spring","summer"}
}

func branch() {
	seasons:= getSeasons()
	for i:=0,i<len(seasons),i++ {
		switch s := ; s {
		case "Spades":
			// ...
		case "Hearts":
			// ...
		case "Diamonds":
			// ...
		case "Clubs":
			// ...
		default:
			panic(fmt.Sprintf("invalid suit %v", s))
		}
	}

}
