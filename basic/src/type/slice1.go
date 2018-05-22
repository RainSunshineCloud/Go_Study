package main

import (
	"fmt"
	"reflect"
)

func main() {
	//方式1
	slice1 := []int{1, 2, 3}
	fmt.Println(slice1)
	//方式2
	slice2 := make([]int, 10)
	for key := range slice2 {
		slice2[key] = 1
	}
	fmt.Println(slice2)
	//方式3
	arr := [5]int{1, 2, 3}
	slice3 := arr[3:]

	fmt.Println(reflect.TypeOf(slice3), slice3)
	//方式4(没啥用)
	slice4 := new([]int)
	fmt.Println(slice4, *slice4, cap(*slice4))
}
