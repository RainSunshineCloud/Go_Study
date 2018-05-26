package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "qwe"
	int := 11
	str1 := string(int)       //表示字符码所在的位置。
	str2 := strconv.Itoa(int) //转化为str11
	//int1 := int(str) //字符串无法通过int() 转化
	if int1, err := strconv.Atoi(str); err != nil {
		fmt.Println(int1)
	} else {
		fmt.Println(err)
	}

	fmt.Println(str, str1, int, str2)
}
