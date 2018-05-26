package main

import (
	"container/list"
	"fmt"
)

func main() {
	lists := list.New()
	lists.PushBack(101)
	lists.PushBack(102)
	lists.PushBack(103)
	fmt.Println(lists.Back().Next())
	for e := lists.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println(lists.Len())
}
