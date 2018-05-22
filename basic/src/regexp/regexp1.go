package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "qdf23232s"
	bytes := []byte(str)
	if ok, _ := regexp.Match("[0-9]+", bytes); ok {
		fmt.Println(str)
	}

	reg, _ := regexp.Compile("[0-9]+")
	byt := reg.ReplaceAll(bytes, []byte("sdf"))
	fmt.Println(string(byt))
}
