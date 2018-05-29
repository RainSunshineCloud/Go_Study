package main

import (
	"fmt"
	"time"
)

/*
time.After 定时器（只运行一次）
*/
func main() {
	ch1 := make(chan string)
	ch2 := make(chan int)
	ch3 := time.After(1e9)
	go test1(ch1)
	go test2(ch2)
	for {
		select {
		case u := <-ch1:
			fmt.Println(u)
		case v := <-ch2:
			fmt.Println(v)
			goto L
		case s := <-ch3:
			fmt.Println(s)
		}
	}
L:
	fmt.Println(2)

}

func test1(ch chan string) {

	ch <- "sdfsk"
}

func test2(ch chan int) {
	time.Sleep(9e9)
	ch <- 1
}
