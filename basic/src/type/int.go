package main

import (
    . "fmt"
)
//全局变量
var int1 = 1
var int2,int3 = 1,1
func main () {
    //局部变量
    int1,int2 := 1,200000000000000000
    Println(int1,int2,int3);

    //整数类型 byte int uint rune int16 uint16 int32 uint32 int64 uint64 uintptr
    //其中有uintptr 存放指针 ；byte 存放字节 rune 存放unicode

    //uintptr 只有在底层交互时才需要（暂时不懂，在声明指针时不需要）
    var bytes byte = 's' //字符编码
    var runes rune = 's' //字符编码
    var p =  &runes
    Println(bytes)
    Println(runes)
    Println(p)

    floats32 := 123123213.232
    floats64 := 21324324324.91321
    //转换byte 最大为64（虽然是int8)
    //超出rune 的范围后转化为一个负数
    //int 转化为float32和float64可转化为科学计数
    //float 转化为int 均为向下取整

    var byte_to_int = int(bytes);
    var int_to_byte = byte(int2);
    var int_to_rune = rune(int2);
    var rune_to_int = int(runes);
    var int_to_float32 = float32(int2);
    var float32_to_int = int(floats32);
    var int_to_float64 = float64(int2);
    var float64_to_int = int(floats64);

    Println("byte_to_int",byte_to_int,"int_to_byte",int_to_byte,"rune_to_int",rune_to_int,"int_to_rune",int_to_rune)
    Println("int_to_float32",int_to_float32,"float32_to_int",float32_to_int,"int_to_float64",int_to_float64,"float64_to_int",float64_to_int)

}


