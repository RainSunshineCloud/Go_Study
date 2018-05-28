package main

import "fmt"
//函数内部若未声明该变量，则会自动找寻全局变量。
//函数内部变量只在函数内部起作用
//

var b int = 1;
func test1() {
	b := 2;
	fmt.Printf("%d",b);
}

func test2() {
	fmt.Printf("%d",b);
}

func main () {
	test2();
	test1();
	test2();
}