package main

import (
	"fmt"
)

//Future
type Function func(string) (string, error)

//Future接口
type Future interface {
	SuccessCallback() error
	FailCallback() error
	Execute(Function) (bool, chan struct{})
}
type UserCache struct {
	Name string
}

func (a *UserCache) SuccessCallback() error {
	fmt.Println("成功回调～")
	return nil
}

func (a *UserCache) FailCallback() error {
	fmt.Println("失败回调～")
	return nil
}

func (a *UserCache) Execute(f Function) (bool, chan struct{}) {
	done := make(chan struct{})
	go func(a *UserCache) {
		_, err := f(a.Name)
		if err != nil {
			_ = a.FailCallback()
		} else {
			_ = a.SuccessCallback()
		}
		done <- struct{}{}
	}(a)
	return true, done
}
func NewUserCache(name string) *UserCache {
	return &UserCache{
		name,
	}
}
func testFuture() {
	var future Future
	future = NewUserCache("Tom")
	updateFunc := func(name string) (string, error) {
		fmt.Println("cache update:", name)
		return name, nil
	}
	_, done := future.Execute(updateFunc)
	defer func() {
		<-done
	}()
}
func main() {
	var future Future
	future = NewUserCache("Barry")
	updateFunc := func(name string) (string, error) {
		fmt.Println("缓存更新：", name)
		return name, nil
	}
	_, done := future.Execute(updateFunc)
	defer func() {
		<-done
	}()
	//...逻辑代码
}

//缓存更新： Barry
//成功回调～
