package main

import (
	"fmt"
	"time"
)

/*
很明显通道会阻塞携程，所以会有通道在等待阻塞中，这时候资源被谁领用就不知道了
根据这种情况资源只会被领用，这相当于消息队列。可以很好的解决抢占问题（重点是原子性）
*/
func main() {
	ch1 := make(chan string)
	go test1(ch1)
	go test2(ch1)
	go test2(ch1)
	time.Sleep(1e10)
}

func test1(ch chan string) {
	time.Sleep(1e9)
	ch <- "dsfs"
}

func test2(ch chan string) {

	fmt.Println(<-ch)
	fmt.Println("exit")
}
