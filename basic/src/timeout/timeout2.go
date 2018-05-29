package main

import (
	"fmt"
	"time"
)

/*
计时器主要有两种：
1.time.NewTicker(d) 需要close (虽然测试过没关闭也没什么影响，估计会占用资源）
2.time.Tick(d) 不需要close
*/
func main() {
	var d time.Duration = 2e9
	ch1 := make(chan string)
	ch2 := make(chan int)
	ch3 := time.Tick(d)
	go test1(ch1)
	go test2(ch2)
	for {
		select {
		case u := <-ch1:
			fmt.Println(u)
		case v := <-ch2:
			fmt.Println(v)
		case s := <-ch3: //周期性跟js 的 interval差不多
			fmt.Println("tick时间到", s)
		}
	}
}

func test1(ch chan string) {
	time.Sleep(5e9)
	ch <- "test1"
}

func test2(ch chan int) {
	time.Sleep(2e9)
	ch <- 1
}
