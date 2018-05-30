package main

import (
	"../iteracter"
	"fmt"
)

func main() {
	var b = [2]int{0, 1}
	ress := iteracter.Factory(f1, b)
	for i := 0; i < 9; i++ {
		res := ress().([2]int)
		fmt.Println(res[1])
	}
}

func f1(b iteracter.Any) iteracter.Any {
	res := b.([2]int)
	res[0], res[1] = res[1], res[0]+res[1]
	return res
}
