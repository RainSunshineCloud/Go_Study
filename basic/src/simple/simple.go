package main

import "fmt"

type Simpler interface {
	Get() int
	Set(int)
}

type Simple struct {
	ints int
}

func main() {
	s := new(Simple)
	var ser Simpler
	ser = s
	ser.Set(1)
	fmt.Println(ser.Get())
}

func (s *Simple) Get() int {
	return s.ints
}

func (s *Simple) Set(ints int) {
	s.ints = ints
}
