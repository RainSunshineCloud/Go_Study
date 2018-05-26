package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a = 1
	pointer := unsafe.Alignof(a)
	size := unsafe.Sizeof(a)
	fmt.Println(pointer, size)
}
