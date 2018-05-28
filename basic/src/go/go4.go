package main

import "fmt"

var sign chan string = make(chan string)

func main() {
	ch := make(chan int)
	go printMsg(ch)
	go sendMsg(ch)
	<-sign

}
func sendMsg(ch chan int) {
	for i := 1; i < 10; i++ {
		ch <- i * 10
	}

}

func printMsg(ch chan int) {
	for i := 1; i < 10; i++ {
		fmt.Println(<-ch)
	}
	go signal()
}

func signal() {
	sign <- "sdf"
}
