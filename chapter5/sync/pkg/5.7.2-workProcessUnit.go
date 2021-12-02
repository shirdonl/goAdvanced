package pkg

import (
	"math/rand"
	"strconv"
	"sync"
	"time"
)

//工作过程
func workProcessUnit(name string, ch chan string) {
	var wd = &WorkProcessData{}

	var group = sync.WaitGroup{}
	group.Add(2)

	go process1(&group, wd)
	go process2(&group, wd)

	group.Wait()

	ch <- name + ":" + wd.GetData()
}

//工作过程1
func process1(group *sync.WaitGroup, gData *WorkProcessData) {
	defer group.Done()
	time.Sleep(time.Microsecond * 1)

	gData.AddData("1", strconv.Itoa(rand.Intn(10)))
}

//工作过程2
func process2(group *sync.WaitGroup, gData *WorkProcessData) {
	defer group.Done()
	time.Sleep(time.Microsecond * 2)

	gData.AddData("2", strconv.Itoa(rand.Intn(10)))
}
