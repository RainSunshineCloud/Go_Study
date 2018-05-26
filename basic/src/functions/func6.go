package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(1)
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}
