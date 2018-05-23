package main

import "fmt"

type tz int

const (
	UTC tz = iota
)

var tzmap = map[tz]string{UTC: "Universal Greenwich time"}

func main() {
	for key, val := range tzmap {
		fmt.Println(key, val)
	}
}
