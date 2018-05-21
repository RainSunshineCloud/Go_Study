package main

import (
	"fmt"
	"strconv"
)

func main() {
	int1, err1 := strconv.Atoi("-10") //字符串数字转化数字
	str1 := strconv.Itoa(1)           //数字转化为字符串
	int2, err2 := strconv.ParseInt("0127", 0, 16)
	//将8进制字符串转化为10进制 第二个参数为0则表示由字符串自身决定是几进制,
	//第三个参数表示转化为int8.int16,int32,0为系统默认int值
	fmt.Println(int1, str1, err1, int2, err2)

}
