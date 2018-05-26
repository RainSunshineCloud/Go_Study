package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bt, _ := ioutil.ReadFile("./file.go") //读取文件
	fmt.Println(string(bt))
}
