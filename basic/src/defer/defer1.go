package main

import "fmt"

/*
在函数执行结束后，就执行defer
注意点： label 对defer没有限制作用，即label 仍会执行。
*/
func main() {
	defer func() {
		fmt.Println("defer 1")
	}()

	test()

	defer func() {
		fmt.Println("defer 3")
	}()

	fmt.Println("end")

}

func test() {
	defer fmt.Println("defer 2")
	if 1 == 1 {

		goto Label1
	} else {
		fmt.Println("defer 7")
		defer fmt.Println("defer 7")
		goto Label2
	}

Label1:
	defer fmt.Println("defer 4")
Label2:
	defer fmt.Println("defer 6")

}
