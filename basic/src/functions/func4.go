package main

import "fmt"

func main() {
	test(9)
}

func test(n uint) {
	fmt.Println(n)
	if n > 1 {
		test(n - 1)
	}
}
