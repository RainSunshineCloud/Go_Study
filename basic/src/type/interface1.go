package main

import "fmt"

type sharper interface {
	area() float32
}

type circle struct {
	r float32
}

//func (circle circle) area() float32 {
//	return 2 * math.Pi * circle.r
//}
func main() {
	//receiver 值和指针都可以，会自动转换
	cs := new(circle)
	cs.r = 1

	var shape sharper
	shape = cs //通过赋值来进行接口的实现

	//cs1是值，不是指针，所以receiver 必须是值
	cs1 := circle{1}
	shape = cs1 //通过赋值来进行接口的实现

	fmt.Println(shape.area())
}
