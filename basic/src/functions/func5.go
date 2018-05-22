package main

import "fmt"

func main() {
	fmt.Println(test(30))
}

func test(n uint64) uint64 {
	if n == 0 {
		return 1
	}

	return n * test(n-1)
}
