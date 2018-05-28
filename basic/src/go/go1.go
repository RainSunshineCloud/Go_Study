package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	//sendMsg(ch) //此时程序死锁，这是因为当前协程阻塞，如果移至go 后面则不会
	//printMsg(ch) //此时程序死锁，系统报错
	//go sendMsg(ch) //阻塞另外一个携程,系统不会报错
	go printMsg(ch) //阻塞另外一个携程，系统不会报错
	sendMsg(ch)     //阻塞当前携程，直到go printMsg接收
	time.Sleep(2e9)
}

func sendMsg(ch chan string) {
	time.Sleep(1e9)
	ch <- "hello world"
	ch <- "你好，世界"
}

func printMsg(ch chan string) {
	var msg string
	for {
		msg = <-ch
		fmt.Println(msg)

	}

}
