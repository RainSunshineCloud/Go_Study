package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Info struct {
	mu  sync.Mutex
	str string
}

func main() {
	var ptr *int32
	var str int32 = 0
	ptr = &str
	bools := atomic.CompareAndSwapInt32(ptr, 0, 1)
	fmt.Println(bools, *ptr)

	var info Info
	info.mu.Lock()
	defer info.mu.Unlock()
	info.str = "sdfsd"
	fmt.Println(info.str)
}
