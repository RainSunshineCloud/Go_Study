package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//scan()
	//scanf()
	//scanln()
	//sscan()
	//sscanf()
	//sscanln()
	bufios()
}

func scan() {
	var a, b int
	int, err := fmt.Scan(&a, &b) //输入间断符为空格
	fmt.Println(int, err, a, b)
}

func scanf() {
	var a, b int
	int, err := fmt.Scanf("%d/%d", &a, &b) //带输入格式
	fmt.Println(int, err, a, b)
}

func scanln() {
	var a, b int
	int, err := fmt.Scanln(&a, &b) //输入间断符为空格
	fmt.Println(int, err, a, b)
}

func sscan() {
	var a = "1111"
	var b, c int
	int, err := fmt.Sscan(a, &b, &c) //与scan一样,但是是从字符串获取
	fmt.Println(int, err, b, c)
}

func sscanf() {
	var a = "112,111"
	var b, c int
	int, err := fmt.Sscanf(a, "%d,%d", &b, &c) //与scanf一样，但是是从字符串获取
	fmt.Println(int, err, b, c)
}

func sscanln() {
	var a = "1111 111"
	var b, c int
	int, err := fmt.Sscanln(a, &b, &c) //与scanln一样，但是是从字符串获取
	fmt.Println(int, err, b, c)
}

func bufios() {

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	fmt.Println(input, err)
}
