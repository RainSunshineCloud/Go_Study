package main

import "fmt"

type stack []interface{}

func (s stack) Pop() stack {
	if len(s) == 0 {
		return s
	}
	news := make(stack, len(s)-1, len(s)-1)
	copy(news, s)

	return news
}

func (s stack) Push(data interface{}) stack {
	news := append(s, data)
	return news
}

func NewStack() stack {
	s := make(stack, 0, 0)
	return s
}

func main() {
	data := map[int]string{1: "1"}
	data1 := map[string]string{"1": "2"}
	stack := NewStack().Push(data).Push(data1).Pop()
	fmt.Println(stack)
}
