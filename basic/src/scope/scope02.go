package main
//函数内部更改全局变量。（和php global 全局变量一致）
import "fmt"

var a int = 1;

func test1 () {
	a = 2;
	fmt.Println(a);
}

func test2 () {
	var a = 1;
	fmt.Println(a);
}

func main () {
	test1();
	test2();
	test1();
}
