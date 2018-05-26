package main

import "fmt"

var i3 *int

func main() {
	i1 := 5
	i2 := &i1
	i3 = &i1
	fmt.Printf("该变量内存地址为：%v\n", &i1)
	fmt.Printf("该变量内存地址为：%v\n", i2)
	fmt.Printf("值为：%v\n", *i2)
	fmt.Printf("该变量内存地址为：%v\n", i3)
}
