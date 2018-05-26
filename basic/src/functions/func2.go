package main

import (
	"fmt"
	"reflect"
)

func main() {
	test(1, 2, 3, 4, 5, 6)
}

func test(s ...int) {
	fmt.Println(s, reflect.TypeOf(s))
	for _, val := range s {
		fmt.Println(val)
	}
}
