package main

import (
	"fmt"
	"time"
)

func main() {
	var d time.Duration = 1e9
	tick := time.NewTicker(d)
	defer tick.Stop()

	ch1 := make(chan string)
	ch2 := make(chan int)
	go test1(ch1)
	go test2(ch2)
	for {
		select {
		case u := <-ch1:
			fmt.Println(u)
		case v := <-ch2:
			fmt.Println(v)
		case s := <-tick.C: //周期性跟js 的 interval差不多
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
