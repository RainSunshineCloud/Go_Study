package main

import (
    . "fmt"
)

func main () {
    str := "sdfsd"
    //转化为byte类型的切片
    bytes := []byte(str)
    Println(str,bytes)

    //修改slice,并转化为string
    bytes[0] = 'f'
    str = string(bytes)
    Println(str)

    //转化为rune类型的切片
    runes := []rune(str)
    Println(str,runes)

    //修改slice,并转化为string
    runes[0] = 'r'
    str = string(runes)
    Println(str)

    //转化为数字(无法转化）
    //ints := []int(str)
    //Println(str,ints)
}