package main

import "fmt"
//虽然同样都是int但因为类型不同，不能相加
//常量也是一样
var int2 int32 = 23;
const cint1 int32 = 23;
const cint2 float32 = 23.3;
const cint3 float64 = 23.2
func main () {
	int1 := 1;
	//fmt.Println(int1 + cint1); //变量和常量（不同类型）报错
	fmt.Println(int1 + int1);//同类型可以运算
	fmt.Println(cint1 + cint1);//同类型可以运算
	fmt.Println(cint1 + 23); //可以参加运算
	//fmt.Println(cint1 + 23.2);//不同类型无法运算
	fmt.Println(cint2 + 23);//可以运算23被当做float32
	fmt.Println(cint3 + 23);//可以运算23被当做float64
	//fmt.Println(int1+int2); //变量和变量相加报错（不同类型）
}