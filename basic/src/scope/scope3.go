package main

import "fmt"
/*
 大括号隔断了作用域(并且将该语句包含的小括号也包含了)
这个与php 很不一样，php 只有函数是局部变量。而main只要是{}就是局部变量
同时，php 没有向上继承这个隐式操作，必须要用global来声明全局，这个部分和js比较像
*/
func main () {
	var res int;
	for i := 1; i < 10 ;i++ {
		res = i;
		fmt.Println(res,i);
	}
	{
		res := 2;
		fmt.Println(res);
	}
	fmt.Println(res);
	//fmt.Println(res,i);//报错，未定义
}