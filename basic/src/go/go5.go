package main

import "fmt"

func main () {

	ch := make(chan int);
	data := []int {1,2,3,4,5,6,7,8,9}
	go sendMsg(ch , data);
	for {
		v,ok := <-ch;
		if ok {
			fmt.Println(v);
		}
	}

}

func sendMsg (ch chan int,data []int) {
	defer close(ch);
	for _,v := range data {
		ch <-v;
	}

}