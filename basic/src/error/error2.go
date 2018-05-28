package main

import (
	"fmt"
)

func main() {
	fmt.Println(divide(2, 0))
	fmt.Println("wer")
	fmt.Println(divide(2, 1))
}

type errors interface{}

func divide(a int, b int) (int, err errors) {
	defer func() {
		if err = recover(); err != nil {
			err = fmt.Errorf("err:%v", err)
		}
	}()

	if b == 0 {
		panic("0 不能作为分母")
	}

	return a / b, nil
}
