package main

import (
	"fmt"
	"time"
)

func main() {
	go test1()
	go test2()
	time.Sleep(6e9)
}

func test1() {
	fmt.Println("sf")
}

func test2() {
	defer func() {
		if errs := recover(); errs != nil {
			fmt.Println(errs)
		}
		a := 1 + 1
		fmt.Println(a)
	}()

	if true {
		panic("sdf")
		fmt.Println("q2")
	}
}
