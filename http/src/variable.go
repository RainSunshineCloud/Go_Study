package main

import (
	"fmt"
	"unsafe"
)

//变量的声明
var var1 int
var var2, var3 int
var var4 int = 1
var var6, var7 int = 1, 1
var var8, var9 = 1, 1

func get() (int, int) {
	//只能在函数内声明
	varn, varm := 1, 1
	return varn, varm
}

//特殊变量_舍去不要的值，这个非常常用，因为golang 默认每个值都要被接收。
var var13, _ = test()

func test() (a int, b int) {
	return a, b
}

//定义单行字符串（因为golang换行代表语句结束
var str = "sdfk"

//多行字符串
var str1 = `sdf
sdfs`

func main() {
	//修改字符串
	var c = []byte(str)
	c[0] = 'c'
	var s2 = string(c)
	fmt.Println(s2)
	//string 无法做 int 类型转换
	//i := []int(str)
	//fmt.Println(i[0])
	//float32 占4个字节，float 占8个字节,int 8个字节 byte 一个字节
	var f float64 = 2343
	i := int(f)
	var strings string = "sdf"
	var bytes byte = 's'
	fmt.Println(unsafe.Sizeof(f))
	fmt.Println(unsafe.Sizeof(i))
	fmt.Println(unsafe.Sizeof(strings))
	fmt.Println(unsafe.Sizeof(bytes))
}
